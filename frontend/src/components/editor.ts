import { computed, ComputedRef, nextTick, Ref, ref, toRaw, watch } from "vue";
import {
    Node,
    Edge,
    useVueFlow,
    NodeSelectionChange,
    MarkerType,
} from "@vue-flow/core";
import { AddPerson, SaveFile as DoSaveFile } from "../../wailsjs/go/main/App";
import { models } from "../../wailsjs/go/models";
import SaveFile = models.SaveFile;
import People = models.Person;
import Family = models.Family;
import { EventsOn } from "../../wailsjs/runtime";
import Person = models.Person;

export function useEditor() {
    // Data
    const saveFile: Ref<SaveFile | null> = ref(null);
    const previousSaveFile: Ref<SaveFile | null> = ref(null);
    const nodes = ref<Node[]>([]);
    const edges = ref<Edge[]>([]);
    const selectedNodes = ref<Node[]>([]);
    const gridCanvas: Ref<HTMLCanvasElement | null> = ref(null);
    let isSaving: Ref<boolean> = ref(false);

    // Composables
    const { getSelectedNodes, onNodesChange, getViewport, project } =
        useVueFlow();

    // Computed properties
    const hasLoadedSaveFile: ComputedRef<boolean> = computed(() => {
        return !!saveFile.value;
    });

    const selectedNode: ComputedRef<Node | null> = computed(() => {
        return selectedNodes.value.length == 1 ? selectedNodes.value[0] : null;
    });

    const readonlyPeople: ComputedRef<Person[]> = computed(() => {
        return (saveFile.value?.people ?? []).map((person) =>
            structuredClone(toRaw(person)),
        );
    });

    const readonlyFamilies: ComputedRef<Family[]> = computed(() => {
        return (saveFile.value?.families ?? []).map((family) =>
            structuredClone(toRaw(family)),
        );
    });

    // Methods
    function init(htmlCanvasElement: HTMLCanvasElement): void {
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
            loadSaveFile(saveFile);
        });

        EventsOn("onSaveRequested", () => {
            saveSaveFile();
        });
    }

    function loadSaveFile(JSON: SaveFile): void {
        // Setting the saveFile to null first helps the watcher to know not to
        // autosave the savefile once.
        saveFile.value = null;
        saveFile.value = JSON;

        initNodes();
        initEdges();
    }

    function initNodes() {
        nodes.value = [];

        if (!saveFile.value) {
            return;
        }

        nodes.value = [
            ...saveFile.value.people.map((person) => ({
                id: "person-" + person.id,
                type: "person",
                position: { x: person.position_x, y: person.position_y },
                data: person,
            })),
            ...saveFile.value.families.map((family) => ({
                id: "family-" + family.id,
                type: "family",
                position: { x: family.position_x, y: family.position_y },
                data: family,
                origin: [12.5, 12.5],
            })),
        ];
    }

    function initEdges() {
        edges.value = [];

        if (!saveFile.value) {
            return;
        }

        edges.value = [
            ...saveFile.value.families.flatMap((family) => [
                ...(family.person_1_id
                    ? [
                          {
                              id: "family-" + family.id + "-male",
                              type: "smoothstep",
                              source: "person-" + family.person_1_id,
                              target: "family-" + family.id,
                              targetHandle: "left",
                              style: { strokeWidth: 2 },
                          },
                      ]
                    : []),
                ...(family.person_2_id
                    ? [
                          {
                              id: "family-" + family.id + "-female",
                              type: "smoothstep",
                              source: "person-" + family.person_2_id,
                              target: "family-" + family.id,
                              targetHandle: "right",
                              style: { strokeWidth: 2 },
                          },
                      ]
                    : []),
            ]),
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
        if (!gridCanvas.value) {
            return;
        }

        const nodeWidth = 256;
        const nodeHeight = 64;
        const gridSize = 50;

        // screen center coordinates
        const centerScreen = {
            x: gridCanvas.value.getBoundingClientRect().width / 2,
            y: gridCanvas.value.getBoundingClientRect().height / 2,
        };

        // convert to graph coordinates
        const centerGraph = project(centerScreen);

        const position = {
            x:
                Math.round((centerGraph.x - nodeWidth / 2) / gridSize) *
                gridSize,
            y:
                Math.round((centerGraph.y - nodeHeight / 2) / gridSize) *
                gridSize,
        };

        AddPerson(position.x, position.y).then((person) => {
            saveFile.value?.people.push(person);

            nodes.value.push({
                id: "person-" + person.id,
                type: "person",
                position: { x: person.position_x, y: person.position_y },
                data: person,
            });
        });
    }

    const saveSaveFile = debounce(() => {
        isSaving.value = true;

        // Send the possibly updated saveFile to the backend for saving
        DoSaveFile({ ...saveFile.value } as SaveFile).finally(
            () => (isSaving.value = false),
        );
    }, 1000);

    function handleNodesSelectionDrag({ node, event }: any): void {
        getSelectedNodes.value.forEach((selectedNode) => {
            selectedNode.data.position_x = selectedNode.position.x;
            selectedNode.data.position_y = selectedNode.position.y;
        });
    }

    function handlePotentiallyChangedConnections(newValue: SaveFile): void {
        let peopleWithChangedConnections: Person[] = newValue.people.filter(
            (newPerson) => {
                let oldPerson = previousSaveFile.value?.people.find(
                    (oldP) => oldP.id === newPerson.id,
                );
                return oldPerson?.family_id !== newPerson.family_id;
            },
        );

        if (peopleWithChangedConnections.length > 0) {
            initEdges();
        }
    }

    function debounce<T extends (...args: any[]) => any>(
        func: T,
        delay: number = 1000,
    ): (...args: Parameters<T>) => void {
        // Prepare a var that will hold the timeout id, so it can be cleared
        let timer: ReturnType<typeof setTimeout>;

        // Return a function that wraps the closure call in a timeout that
        // clears (resets) every time the debounced function is called before
        // the timer runs out.
        return (...args: Parameters<T>) => {
            clearTimeout(timer);
            timer = setTimeout(() => {
                func(...args);
            }, delay);
        };
    }

    watch(
        () => saveFile.value,
        (newValue, oldValue) => {
            // This check prevents saving immediately after loading a savefile.
            if (newValue !== null && oldValue !== null) {
                // The previous state of the saveFile is tracked manually to
                // work around limitations with 'deep' watchers.
                if (previousSaveFile.value) {
                    handlePotentiallyChangedConnections(newValue);
                }

                saveSaveFile();
            }

            previousSaveFile.value = structuredClone(toRaw(newValue));
        },
        { deep: true },
    );

    return {
        nodes,
        edges,
        hasLoadedSaveFile,
        selectedNode,
        readonlyPeople,
        readonlyFamilies,
        init,
        handleNodesSelectionDrag,
        isSaving,
        addPerson,
    };
}
