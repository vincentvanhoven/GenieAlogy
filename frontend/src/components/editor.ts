import { computed, ComputedRef, nextTick, Ref, ref, toRaw } from "vue";
import {
    Node,
    Edge,
    useVueFlow,
    NodeSelectionChange,
    MarkerType,
} from "@vue-flow/core";
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
    let offset = { x: 0, y: 0, zoom: 1 };

    // Composables
    const { getSelectedNodes, onNodesChange } = useVueFlow();

    // Computed properties
    const selectedNode: ComputedRef<Node | null> = computed(() => {
        return selectedNodes.value.length == 1 ? selectedNodes.value[0] : null;
    });

    // Methods
    function init(htmlCanvasElement: HTMLCanvasElement) {
        gridCanvas.value = htmlCanvasElement;

        // Event listeners
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

        loadConfiguration({
            people: <People[]>[],
            families: <Family[]>[],
        } as SaveFile);
    }

    function loadConfiguration(JSON: SaveFile) {
        saveFile.value = JSON;

        nodes.value = [
            ...saveFile.value.people.map((person) => ({
                id: "person-" + person.uuid,
                type: "person",
                position: { x: person.position_x, y: person.position_y },
                data: person,
            })),
            ...saveFile.value.families.map((family) => ({
                id: "family-" + family.uuid,
                type: "family",
                position: { x: family.position_x, y: family.position_y },
                data: family,
                origin: [12.5, 12.5],
            })),
        ];

        edges.value = [
            ...saveFile.value.families.flatMap((family) => [
                ...(family.person_1_uuid
                    ? [
                          {
                              id: "family-" + family.uuid + "-male",
                              type: "smoothstep",
                              source: "person-" + family.person_1_uuid,
                              target: "family-" + family.uuid,
                              targetHandle: "left",
                              style: { strokeWidth: 2 },
                          },
                      ]
                    : []),
                ...(family.person_2_uuid
                    ? [
                          {
                              id: "family-" + family.uuid + "-female",
                              type: "smoothstep",
                              source: "person-" + family.person_2_uuid,
                              target: "family-" + family.uuid,
                              targetHandle: "right",
                              style: { strokeWidth: 2 },
                          },
                      ]
                    : []),
            ]),
            ...saveFile.value.people
                .filter((person) => person.family_uuid)
                .map((person) => ({
                    id:
                        "family-" +
                        person.family_uuid +
                        "-child-" +
                        person.uuid,
                    type: "smoothstep",
                    source: "family-" + person.family_uuid,
                    target: "person-" + person.uuid,
                    markerEnd: {
                        type: MarkerType.ArrowClosed,
                        color: "black",
                    },
                })),
        ];
    }

    function saveConfiguration() {
        DoSaveFile({ ...saveFile.value } as SaveFile);
    }

    function handleNodesSelectionDrag({ node, event }: any) {
        getSelectedNodes.value.forEach((selectedNode) => {
            console.log(
                selectedNode.data.firstname,
                selectedNode.position.x,
                selectedNode.position.y,
            );

            selectedNode.data.position_x = selectedNode.position.x;
            selectedNode.data.position_y = selectedNode.position.y;
        });
    }

    function onMove({ event, flowTransform }: any) {
        offset.x = flowTransform.x;
        offset.y = flowTransform.y;
        offset.zoom = flowTransform.zoom;
    }

    return {
        nodes,
        edges,
        selectedNode,
        init,
        handleNodesSelectionDrag,
        onMove,
    };
}
