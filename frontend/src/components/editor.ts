import { computed, ComputedRef, nextTick, Ref, ref, toRaw } from "vue";
import { Node, Edge, useVueFlow, NodeSelectionChange } from "@vue-flow/core";
import { SaveFile as DoSaveFile } from "../../wailsjs/go/main/App";
import { models } from "../../wailsjs/go/models";
import SaveFile = models.SaveFile;
import People = models.Person;
import Family = models.Family;
import { EventsOn } from "../../wailsjs/runtime";

export function useEditor() {
    // Data
    const saveFile: Ref<SaveFile | null> = ref(null);
    const nodes = ref<Node[]>([]);
    const edges = ref<Edge[]>([]);
    const selectedNodes = ref<Node[]>([]);
    const gridCanvas: Ref<HTMLCanvasElement | null> = ref(null);

    const baseCellSize = 50;
    const snapGrid = [50, 300];
    const dragStart = new Map<string, { x: number; y: number }>();
    let offset = { x: 0, y: 0, zoom: 1 };
    let initialMouse = { x: 0, y: 0 };
    const nodeWidth = 100;

    // Composables
    const { project, updateNode, getSelectedNodes, onNodesChange } = useVueFlow();

    // Computed properties
    const selectedNode: ComputedRef<Node | null> = computed(() => {
        return selectedNodes.value.length == 1 ? selectedNodes.value[0] : null;
    });

    // Methods
    function init(htmlCanvasElement: HTMLCanvasElement) {
        gridCanvas.value = htmlCanvasElement;

        // Event listeners
        window.addEventListener("resize", drawGrid);

        onNodesChange((changes) => {
            // Update selectedNodes
            changes
                .filter((change) => Object.hasOwn(change, "selected"))
                .map((change) => change as NodeSelectionChange)
                .forEach((change) => {
                    if (change.selected) {
                        selectedNodes.value.push(
                            nodes.value.find(
                                (node) => node.id === change.id,
                            ) as Node,
                        );
                    } else {
                        const index = selectedNodes.value.findIndex(
                            (node) => node.id === change.id,
                        );
                        if (index !== -1) {
                            selectedNodes.value.splice(index, 1);
                        }
                    }
                });
        });

        EventsOn("onSaveFileLoaded", (saveFile: SaveFile) => {
            loadConfiguration(saveFile);
        });

        EventsOn("onSaveRequested", () => {
            saveConfiguration();
        });

        loadConfiguration({people: <People[]>[], families: <Family[]>[]} as SaveFile);
    }

    function loadConfiguration(JSON: SaveFile) {
        saveFile.value = JSON;

        let dereffedSaveFile = { ...saveFile.value };

        nodes.value = [
            ...dereffedSaveFile.people.map((person) => ({
                id: "person-" + person.uuid,
                type: "person",
                position: {x: person.position_x, y: person.position_y},
                data: person,
            })),
            ...dereffedSaveFile.families.map((family) => ({
                id: "family-" + family.uuid,
                type: "family",
                position: { x: 0, y: 0 },
                // draggable: false,
                data: family,
            })),
        ];

        edges.value = [
            ...dereffedSaveFile.families.flatMap((family) => [
                ...(family.person_1_uuid
                    ? [
                          {
                              id: "family-" + family.uuid + "-male",
                              type: "straight",
                              source: "person-" + family.person_1_uuid,
                              target: "family-" + family.uuid,
                              style: { strokeWidth: 2 },
                          },
                      ]
                    : []),
                ...(family.person_2_uuid
                    ? [
                          {
                              id: "family-" + family.uuid + "-female",
                              type: "straight",
                              source: "person-" + family.person_2_uuid,
                              target: "family-" + family.uuid,
                              style: { strokeWidth: 2 },
                          },
                      ]
                    : []),
            ]),
            ...dereffedSaveFile.people
                .filter((person) => person.family_uuid)
                .map((person) => ({
                    id: "family-" + person.family_uuid + "-child-" + person.uuid,
                    type: "smoothstep",
                    source: "family-" + person.family_uuid,
                    target: "person-" + person.uuid,
                })),
        ];

        nextTick(() => {
            adjustAllFamilyNodePositions();
        });
    }

    function saveConfiguration() {
        DoSaveFile({ ...saveFile.value } as SaveFile);
    }

    function drawGrid() {
        if (!gridCanvas.value) {
            return;
        }

        const ctx = gridCanvas.value?.getContext("2d")!;
        const width = gridCanvas.value.clientWidth;
        const height = gridCanvas.value.clientHeight;
        gridCanvas.value.width = width;
        gridCanvas.value.height = height;

        const zoom = offset.zoom;
        const pattern = [3, 3]; // 3 rows filled, 3 rows blank
        const colors = [null, "#f0f0f0"];
        const patternHeight = 6 * baseCellSize;

        ctx.clearRect(0, 0, width, height);
        ctx.save();

        ctx.translate(offset.x, offset.y);
        ctx.scale(zoom, zoom);

        const startX = -offset.x / zoom;
        const startY = -offset.y / zoom;
        const endX = startX + baseCellSize + width / zoom;
        const endY = startY + height / zoom;

        // Background stripes
        let yWorld = Math.floor(startY / patternHeight) * patternHeight;

        while (yWorld < endY) {
            let stripeY = yWorld;

            for (let i = 0; i < pattern.length; i++) {
                const rows = pattern[i];
                const stripeHeight = rows * baseCellSize;
                const color = colors[i];

                if (color && stripeY + stripeHeight > startY) {
                    ctx.fillStyle = color;
                    ctx.fillRect(
                        Math.floor(startX / baseCellSize) * baseCellSize,
                        stripeY,
                        Math.ceil((endX - startX) / baseCellSize) *
                            baseCellSize,
                        stripeHeight,
                    );
                }

                stripeY += stripeHeight;
                if (stripeY > endY) break;
            }

            yWorld += patternHeight;
        }

        // Vertical grid lines (only in first 3 rows of pattern)
        ctx.strokeStyle = "#aaa";

        for (
            let x = Math.floor(startX / baseCellSize) * baseCellSize;
            x <= endX;
            x += baseCellSize
        ) {
            let yWorld = Math.floor(startY / baseCellSize) * baseCellSize;

            while (yWorld <= endY) {
                const yInPattern =
                    ((yWorld % patternHeight) + patternHeight) % patternHeight;

                if (yInPattern < 3 * baseCellSize) {
                    ctx.beginPath();
                    ctx.moveTo(x, yWorld);
                    ctx.lineTo(x, Math.min(yWorld + baseCellSize, endY));
                    ctx.stroke();
                }

                yWorld += baseCellSize;
            }
        }

        // Horizontal grid lines (only in first 3 rows)
        yWorld = Math.floor(startY / baseCellSize) * baseCellSize;
        while (yWorld <= endY) {
            const yInPattern =
                ((yWorld % patternHeight) + patternHeight) % patternHeight;

            if (yInPattern <= 3 * baseCellSize) {
                ctx.beginPath();
                ctx.moveTo(startX, yWorld);
                ctx.lineTo(endX, yWorld);
                ctx.stroke();
            }

            yWorld += baseCellSize;
        }

        ctx.restore();
    }

    function handleNodesSelectionDragStart({ node, event }: any) {
        const graphPos = project({ x: event.clientX, y: event.clientY });
        initialMouse = { x: graphPos.x, y: graphPos.y };
        dragStart.clear();

        getSelectedNodes.value.forEach((selectedNode) => {
            dragStart.set(selectedNode.id, {
                x: selectedNode.position.x,
                y: selectedNode.position.y,
            });
        });
    }

    function handleNodesSelectionDrag({ node, event }: any) {
        const graphPos = project({ x: event.clientX, y: event.clientY });
        const dx = graphPos.x - initialMouse.x;
        const dy = graphPos.y - initialMouse.y;

        getSelectedNodes.value
            .filter((selectedNode) => selectedNode.type !== "family")
            .forEach((selectedNode) => {
                handleNodeDrag(selectedNode, dx, dy);
            });
    }

    function handleNodeDrag(node: any, deltaX: number, deltaY: number) {
        const start = dragStart.get(node.id)!;
        const newPosition = { x: start.x + deltaX, y: start.y + deltaY };

        if (node.type === "person") {
            updateNodePosition(node, {
                x: Math.round(newPosition.x / snapGrid[0]) * snapGrid[0],
                y: Math.round(newPosition.y / snapGrid[1]) * snapGrid[1],
            });

            adjustFamilyNodePositions(node);
        }
    }

    function adjustAllFamilyNodePositions() {
        nodes.value.forEach((node) => adjustFamilyNodePositions(node));
    }

    function adjustFamilyNodePositions(spouseNode: Node) {
        nodes.value
            .filter((familyNode) => familyNode.type === "family")
            .filter((familyNode) => {
                return [
                    `person-${familyNode.data.person_1_id}`,
                    `person-${familyNode.data.person_2_id}`,
                ].includes(spouseNode.data.uuid);
            })
            .forEach((familyNode) => {
                let searchNodeId =
                    spouseNode.data.uuid === familyNode.data.person_1_id
                        ? familyNode.data.person_2_id
                        : familyNode.data.person_1_id;

                let partnerNode = nodes.value.find(
                    (partnerNode) =>
                        partnerNode.data.uuid === searchNodeId,
                ) as Node;
                updateNodePosition(partnerNode, { y: spouseNode.position.y });

                let leftPoint = partnerNode.position.x + nodeWidth;
                let rightPoint = spouseNode.position.x;
                updateNodePosition(familyNode, {
                    x: leftPoint - (leftPoint - rightPoint) / 2 - 12.5,
                    y: spouseNode.position.y + 50 + 12.5,
                });

                if (partnerNode.position.x !== spouseNode.position.x) {
                    let leftNode =
                        partnerNode.position.x < spouseNode.position.x
                            ? partnerNode
                            : spouseNode;
                    let rightNode =
                        partnerNode.position.x > spouseNode.position.x
                            ? partnerNode
                            : spouseNode;

                    let maleEdge = edges.value.find(
                        (edge) =>
                            familyNode.id.toString() + "-male" === edge.id,
                    );
                    let femaleEdge = edges.value.find(
                        (edge) =>
                            familyNode.id.toString() + "-female" === edge.id,
                    );

                    if (maleEdge && maleEdge.source == leftNode.id) {
                        maleEdge.sourceHandle = "right";
                        maleEdge.targetHandle = "left";
                    } else if (maleEdge && maleEdge.source == rightNode.id) {
                        maleEdge.sourceHandle = "left";
                        maleEdge.targetHandle = "right";
                    }

                    if (femaleEdge && femaleEdge.source == leftNode.id) {
                        femaleEdge.sourceHandle = "right";
                        femaleEdge.targetHandle = "left";
                    } else if (
                        femaleEdge &&
                        femaleEdge.source == rightNode.id
                    ) {
                        femaleEdge.sourceHandle = "left";
                        femaleEdge.targetHandle = "right";
                    }
                }
            });
    }

    function updateNodePosition(
        node: Node,
        position: { x?: number; y?: number },
    ) {
        let oldPosition = { ...node.position };
        let newPosition = {
            x: position.x ?? node.position.x,
            y: position.y ?? node.position.y,
        };

        if (
            newPosition.x !== oldPosition.x ||
            newPosition.y !== oldPosition.y
        ) {
            updateNode(node.id, (node) => ({
                position: newPosition,
            }));


            // TODO: ref this better
            let person = saveFile.value?.people.find(p => p.uuid == node.data.uuid);

            if(person) {
                person.position_x = newPosition.x;
                person.position_y = newPosition.y;
            }
        }
    }

    function onMove({ event, flowTransform }: any) {
        offset.x = flowTransform.x;
        offset.y = flowTransform.y;
        offset.zoom = flowTransform.zoom;

        drawGrid();
    }

    return {
        nodes,
        edges,
        selectedNode,
        init,
        drawGrid,
        handleNodesSelectionDragStart,
        handleNodesSelectionDrag,
        onMove,
    };
}
