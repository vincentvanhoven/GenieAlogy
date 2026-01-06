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
                @move="editor.onMove"
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
        </div>

        <Sidebar :selected-node="editor.selectedNode.value" />
    </div>
</template>

<script lang="ts" setup>
    import { VueFlow } from "@vue-flow/core";
    import PersonNode from "./PersonNode.vue";
    import FamilyNode from "./FamilyNode.vue";
    import { onMounted, Ref, ref } from "vue";
    import { useEditor } from "./editor";
    import Sidebar from "./Sidebar.vue";
    import { Background } from "@vue-flow/background";

    const editor = useEditor();
    const gridCanvas: Ref<HTMLCanvasElement | null> = ref(null);

    // Event listeners
    onMounted(() => editor.init(gridCanvas.value!));
</script>
