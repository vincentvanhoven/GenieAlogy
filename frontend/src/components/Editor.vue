<template>
    <div class="w-full h-full grid grid-cols-4">
        <div class="relative w-full h-full col-span-3">
            <canvas
                ref="gridCanvas"
                class="absolute top-0 left-0 w-full h-full pointer-events-none"
            />

            <VueFlow
                class="relative w-full h-full"
                v-model:nodes="editor.nodes.value"
                v-model:edges="editor.edges.value"
                @node-drag="editor.handleNodesSelectionDrag"
                :snap-grid="[16, 16]"
                :snap-to-grid="true"
                :min-zoom="0.01"
                :max-zoom="1"
            >
                <template #node-person="props">
                    <PersonNode v-bind="props" />
                </template>

                <template #node-family="props">
                    <FamilyNode v-bind="props" />
                </template>

                <Background variant="lines" :gap="16" />
            </VueFlow>

            <div
                v-show="editor.hasLoadedSaveFile.value"
                class="absolute right-3 top-2 flex gap-2 rounded shadow p-2
                    bg-white border border-gray-200"
            >
                <button
                    @click="editor.addPerson"
                    type="button"
                    class="rounded border border-gray-500 p-0.5 cursor-pointer
                        hover:border-gray-400 hover:bg-gray-200 hover:shadow
                        active:bg-gray-300"
                >
                    <img src="../assets/images/plus.svg" class="w-4 h-4" />
                </button>

                <button
                    type="button"
                    @click="editor.removeSelectedPerson"
                    :disabled="!editor.selectedNode.value"
                    class="rounded border border-gray-500 p-0.5 cursor-pointer
                        hover:border-gray-400 hover:bg-gray-200 hover:shadow
                        active:bg-gray-300
                        disabled:cursor-not-allowed disabled:bg-gray-200
                        disabled:hover:border-gray-500 disabled:hover:shadow-none"
                >
                    <img
                        src="../assets/images/trash.svg"
                        class="w-4 h-4 p-0.5"
                    />
                </button>
            </div>

            <div
                class="absolute left-3 bottom-2 transition-all duration-300"
                :class="!!showSavingMessage ? 'opacity-100' : 'opacity-0'"
            >
                {{ savingMessage }}
            </div>
        </div>

        <Sidebar
            :people="editor.readonlyPeople.value"
            :families="editor.readonlyFamilies.value"
            :selected-node="editor.selectedNode.value"
        />
    </div>
</template>

<script lang="ts" setup>
    import { VueFlow } from "@vue-flow/core";
    import PersonNode from "./PersonNode.vue";
    import FamilyNode from "./FamilyNode.vue";
    import { onMounted, Ref, ref, watch } from "vue";
    import { useEditor } from "./editor";
    import Sidebar from "./Sidebar.vue";
    import { Background } from "@vue-flow/background";

    // Data
    const editor = useEditor();
    const gridCanvas: Ref<HTMLCanvasElement | null> = ref(null);
    const showSavingMessage: Ref<boolean> = ref(false);
    const savingMessage: Ref<string | null> = ref(null);
    const clearMessageTimeout: Ref<number | null> = ref(null);

    // Event listeners
    onMounted(() => editor.init(gridCanvas.value!));

    // Watchers
    watch(editor.isSaving, (value, oldValue) => {
        if (clearMessageTimeout.value) {
            clearTimeout(clearMessageTimeout.value);
            clearMessageTimeout.value = null;
        }

        // If currently saving
        if (value) {
            savingMessage.value = "Autosaving...";
            showSavingMessage.value = true;
        } else {
            // If there was an update, and not currently saving, saving must have finished.
            savingMessage.value += " done!";
            clearMessageTimeout.value = window.setTimeout(
                () => (showSavingMessage.value = false),
                1500,
            );
        }
    });
</script>
