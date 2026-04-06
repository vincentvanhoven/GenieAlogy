<template>
    <div class="w-screen h-screen overflow-hidden">
        <div class="w-full h-full grid grid-cols-4">
            <div class="relative w-full h-full col-span-3">
                <canvas
                    ref="gridCanvas"
                    class="absolute top-0 left-0 w-full h-full
                        pointer-events-none"
                />

                <VueFlow
                    :title="saveFileStore.isEditingPerson ? 'Please finish editing in the sidebar first' : ''"
                    :class="{'cursor-not-allowed' : saveFileStore.isEditingPerson }"
                    class="relative w-full h-full"
                    v-model:nodes="saveFileStore.nodes"
                    v-model:edges="saveFileStore.edges"
                    @node-drag="saveFileStore.handleNodesSelectionDrag"

                    :elements-selectable="!saveFileStore.isEditingPerson"
                    :nodes-draggable="!saveFileStore.isEditingPerson"
                    :select-nodes-on-drag="!saveFileStore.isEditingPerson"

                    :nodes-connectable="false"
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
                    v-show="saveFileStore.hasLoadedSaveFile"
                    class="absolute right-3 top-2 flex gap-2 rounded shadow p-2
                        bg-white border border-gray-200"
                >
                    <button
                        @click="saveFileStore.addPerson"
                        type="button"
                        class="rounded border border-gray-500 p-0.5
                            cursor-pointer hover:border-gray-400
                            hover:bg-gray-200 hover:shadow active:bg-gray-300"
                    >
                        <img src="./assets/images/plus.svg" class="w-4 h-4" />
                    </button>
                </div>

                <div
                    class="absolute left-3 bottom-2 transition-all duration-300"
                    :class="!!showSavingMessage ? 'opacity-100' : 'opacity-0'"
                >
                    {{ savingMessage }}
                </div>
            </div>

            <Sidebar />
        </div>
    </div>
</template>

<script lang="ts" setup>
    import { VueFlow } from "@vue-flow/core";
    import PersonNode from "./components/PersonNode.vue";
    import FamilyNode from "./components/FamilyNode.vue";
    import { onMounted, Ref, ref, watch } from "vue";
    import Sidebar from "./components/Sidebar.vue";
    import { Background } from "@vue-flow/background";
    import { useSaveFileStore } from "./stores/saveFileStore";

    // Data
    const saveFileStore = useSaveFileStore();
    const gridCanvas: Ref<HTMLCanvasElement | null> = ref(null);
    const showSavingMessage: Ref<boolean> = ref(false);
    const savingMessage: Ref<string | null> = ref(null);
    const clearMessageTimeout: Ref<number | null> = ref(null);

    // Event listeners
    onMounted(() => saveFileStore.init(gridCanvas.value!));

    // Watchers
    watch(
        () => saveFileStore.isSaving,
        (value, oldValue) => {
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
        },
    );
</script>
