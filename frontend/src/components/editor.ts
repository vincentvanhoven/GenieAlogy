import { computed, ComputedRef, nextTick, Ref, ref, toRaw, watch } from "vue";
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
import Person = models.Person;

export function useEditor() {
    // Data
    const saveFile: Ref<SaveFile | null> = ref(null);
    const nodes = ref<Node[]>([]);
    const edges = ref<Edge[]>([]);
    const selectedNodes = ref<Node[]>([]);
    const gridCanvas: Ref<HTMLCanvasElement | null> = ref(null);
    let isSaving: Ref<boolean> = ref(false);

    // Composables
    const { getSelectedNodes, onNodesChange } = useVueFlow();

    // Computed properties
    const selectedNode: ComputedRef<Node | null> = computed(() => {
        return selectedNodes.value.length == 1 ? selectedNodes.value[0] : null;
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
        (value, oldValue) => {
            // This check prevents saving immediately after loading a savefile.
            if (value !== null && oldValue !== null) {
                saveSaveFile();
            }
        },
        { deep: true },
    );

    return {
        nodes,
        edges,
        selectedNode,
        init,
        handleNodesSelectionDrag,
        isSaving,
    };
}
