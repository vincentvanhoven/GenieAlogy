import { defineStore } from "pinia";
import {
    computed,
    ComputedRef,
    readonly,
    ref,
    Ref,
    WritableComputedRef,
} from "vue";
import {
    Edge,
    GraphNode,
    MarkerType,
    Node,
    NodeSelectionChange,
    useVueFlow,
} from "@vue-flow/core";
import {
    AddFamily,
    AddPerson,
    RemoveFamily,
    RemovePerson,
    UpdatePerson,
    UpdateFamily,
} from "../../wailsjs/go/main/App";
import { models } from "../../wailsjs/go/models";
import { EventsOn } from "../../wailsjs/runtime";
import SaveFile = models.SaveFile;
import Family = models.Family;
import Person = models.Person;

export const useSaveFileStore = defineStore("saveFile", () => {
    // Consts
    const personNodeDimensions = { x: 256, y: 64 };
    const gridSize = 50;

    // References
    const gridCanvas: Ref<HTMLCanvasElement | null> = ref(null);

    // Data
    const saveFile: Ref<SaveFile | null> = ref(null);
    const nodes = ref<Node[]>([]);
    const edges = ref<Edge[]>([]);
    const selectedNodes = ref<Node[]>([]);

    // State flags
    let isSaving: Ref<boolean> = ref(false);
    let isEditingPerson: Ref<boolean> = ref(false);

    // Composables
    const {
        getSelectedNodes,
        addSelectedNodes,
        onNodesChange,
        getViewport,
        project,
    } = useVueFlow();

    // Computed properties
    const hasLoadedSaveFile: ComputedRef<boolean> = computed(() => {
        return !!saveFile.value;
    });

    const selectedNode: ComputedRef<Node | null> = computed(() => {
        return selectedNodes.value.length == 1 ? selectedNodes.value[0] : null;
    });

    const selectedPerson: WritableComputedRef<Person | null> = computed({
        get() {
            return selectedNode.value
                ? getPersonFromNode(selectedNode.value)
                : null;
        },
        set(updateValue: Person) {
            if (selectedNode.value) {
                let person = getPersonFromNode(selectedNode.value);

                if (person) {
                    Object.assign(person, { ...updateValue });
                }
            }
        },
    });

    // Methods
    function init(htmlCanvasElement: HTMLCanvasElement): void {
        gridCanvas.value = htmlCanvasElement;

        // Event listeners
        onNodesChange((changes) => {
            // Update selectedNodes
            let selectChanges = changes.filter((change) =>
                Object.hasOwn(change, "selected"),
            );

            if (selectChanges.length > 0 && !isEditingPerson.value) {
                selectChanges.forEach((change) =>
                    onNodeSelected(change as NodeSelectionChange),
                );
            }
        });

        EventsOn("onSaveFileLoaded", (saveFile: SaveFile) => {
            loadSaveFile(saveFile);
        });

        EventsOn("onSaveRequested", () => {
            // saveSaveFile();
        });
    }

    function onNodeSelected(change: NodeSelectionChange) {
        if (change.selected) {
            // Find the node that was selected
            let selectedNode = nodes.value.find(
                (node) => node.id === change.id,
            ) as Node;
            // Add it to the array of selected nodes
            selectedNodes.value.push(selectedNode);
        } else {
            // Find the node that was deselected
            const index = selectedNodes.value.findIndex(
                (node) => node.id === change.id,
            );
            // Remote it from the array of selected nodes
            if (index !== -1) {
                selectedNodes.value.splice(index, 1);
            }
        }
    }

    function loadSaveFile(JSON: SaveFile): void {
        saveFile.value = JSON;

        // Reset the nodes
        nodes.value = [];
        selectedNodes.value = [];

        // Init the new nodes
        nodes.value = [
            // Person nodes
            ...saveFile.value.people.map((person) => ({
                id: "person-" + person.id,
                type: "person",
                position: { x: person.position_x, y: person.position_y },
            })),
            // Family nodes
            ...saveFile.value.families.map((family) => ({
                id: "family-" + family.id,
                type: "family",
                position: { x: family.position_x, y: family.position_y },
                origin: [12.5, 12.5],
            })),
        ];

        // Reset the edges
        edges.value = [];

        // Init the new edges
        edges.value = [
            // Relationship edges
            ...saveFile.value.families.flatMap((family) => [
                {
                    id: `family-${family.id}-male`,
                    type: "smoothstep",
                    source: "person-" + family.male_id,
                    target: "family-" + family.id,
                    style: { strokeWidth: 2 },
                },
                {
                    id: `family-${family.id}-female`,
                    type: "smoothstep",
                    source: "person-" + family.female_id,
                    target: "family-" + family.id,
                    style: { strokeWidth: 2 },
                },
            ]),
            // Family edges
            ...saveFile.value.people
                .filter((person) => person.family_id)
                .map((person) => ({
                    id: "family-" + person.family_id + "-child-" + person.id,
                    type: "smoothstep",
                    source: "family-" + person.family_id,
                    target: "person-" + person.id,
                    markerEnd: {
                        type: MarkerType.ArrowClosed,
                        color: "black",
                    },
                })),
        ];
    }

    function addPerson() {
        // screen center coordinates
        const centerScreen = {
            x: gridCanvas.value!.getBoundingClientRect().width / 2,
            y: gridCanvas.value!.getBoundingClientRect().height / 2,
        };
        // convert to graph coordinates
        const centerGraph = project(centerScreen);

        const position = {
            x:
                Math.round(
                    (centerGraph.x - personNodeDimensions.x / 2) / gridSize,
                ) * gridSize,
            y:
                Math.round(
                    (centerGraph.y - personNodeDimensions.y / 2) / gridSize,
                ) * gridSize,
        };

        AddPerson(position.x, position.y).then((person) => {
            saveFile.value?.people.push(person);

            nodes.value.push({
                id: "person-" + person.id,
                type: "person",
                position: { x: person.position_x, y: person.position_y },
            });
        });
    }

    function removeSelectedPerson(person: Person) {
        // Trigger removal
        RemovePerson(person).then((newSaveFile: SaveFile) => {
            // TODO: receive only status OK/FAIL and update local state accordingly
        });
    }

    function removeFamily(family: Family) {
        RemoveFamily(family).then((newSaveFile: SaveFile) => {
            // TODO: receive only status OK/FAIL and update local state accordingly
        });
    }

    function addFamily(family: Family) {
        AddFamily(family).then((newSaveFile: SaveFile) => {
            // TODO: receive only status OK/FAIL and update local state accordingly
        });
    }

    function handleNodesSelectionDrag({ node, event }: any): void {
        getSelectedNodes.value.forEach(async (selectedNode) => {
            let person = getPersonFromNode(selectedNode);
            let family = getFamilyFromNode(selectedNode);

            let positionData = {
                position_x: selectedNode.position.x,
                position_y: selectedNode.position.y,
            };

            try {
                if (person) {
                    let updatePayload = { ...person, ...positionData };
                    await UpdatePerson(updatePayload);
                } else if (family) {
                    let updatePayload = { ...family, ...positionData };
                    await UpdateFamily(updatePayload);
                }
            } catch (exception) {
                // console.log(exception);
            }
        });
    }

    function getPersonFromNode(node: Node | string): Person | null {
        // Get the person ID from the Node id
        let id = typeof node === "object" ? node.id : node;
        id = id.replace("person-", "");

        // Assert the result is numerical
        if (id === "" || Number.isNaN(+id)) {
            return null;
        }

        // Find and return the corresponding Person
        return getPersonFromId(+id);
    }

    function getPersonFromId(id: number): Person | null {
        return (
            saveFile.value?.people.find((person) => person.id === +id) ?? null
        );
    }

    function getPersonDisplayName(person: Person) {
        return `${person.firstname ?? "Onbekend"} ${person.lastname ?? ""} (${person.id})`;
    }

    function getFamilyFromNode(node: Node | string): Family | null {
        // Get the family ID from the Node id
        let id = typeof node === "object" ? node.id : node;
        id = id.replace("family-", "");

        // Assert the result is numerical
        if (id === "" || Number.isNaN(+id)) {
            return null;
        }

        // Find and return the corresponding Person
        return getFamilyFromId(+id);
    }

    function getFamilyFromId(id: number): Family | null {
        return (
            saveFile.value?.families.find((family) => family.id === +id) ?? null
        );
    }

    function getFamilyDisplayName(family: Family): string {
        return (
            saveFile.value?.people // Iterate all people
                // Get the parents of this family
                .filter(
                    (person) =>
                        family.male_id === person.id ||
                        family.female_id === person.id,
                )
                // Format their names
                .map((person) => getPersonDisplayName(person))
                // Join their names into one string
                .join(" - ") ?? ""
        );
    }

    function enableEditMode() {
        // Ensure that Person edit mode is not already active, and that there is only one selected node
        if (isEditingPerson.value || !selectedNode.value) {
            return;
        }

        // 'Lock' the graph view for selection changes
        isEditingPerson.value = true;
        // 'Lock' the selected styling of the active Node
        selectedNode.value!.class = "selected";
    }

    async function disableEditMode(changedFields: Person | null = null) {
        // Ensure that Person edit mode is active
        if (!isEditingPerson.value) {
            return;
        }

        if (changedFields) {
            try {
                let updatePayload = {
                    ...selectedPerson.value,
                    ...changedFields,
                };
                await UpdatePerson(updatePayload);
                selectedPerson.value = updatePayload;
            } catch (exception) {
                // console.log(exception);
            }
        }

        // Ensure that VueFlow selected states are in line with the manually
        // forced selection of selectedNode (it may have gotten deselected under
        // the VueFlow hood)
        addSelectedNodes([selectedNode.value! as GraphNode]);

        // 'Unlock' the graph view for selection changes
        isEditingPerson.value = false;
        // 'Unlock' the selected styling of the active Node
        selectedNode.value!.class = "";
    }

    return {
        nodes,
        edges,
        hasLoadedSaveFile,
        selectedNode,
        selectedPerson: readonly(selectedPerson),
        people: computed(() => saveFile.value?.people ?? []),
        families: computed(() => saveFile.value?.families ?? []),
        init,
        handleNodesSelectionDrag,
        isSaving,
        isEditingPerson,
        addPerson,
        removeSelectedPerson,
        addFamily,
        removeFamily,
        getPersonFromNode,
        getPersonFromId,
        getPersonDisplayName,
        getFamilyDisplayName,
        enableEditMode,
        disableEditMode,
    };
});
