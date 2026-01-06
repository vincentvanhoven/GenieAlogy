<template>
    <div
        class="vue-flow__node-default node-root flex justify-center items-center !bg-gray-200 !rounded-full !border !p-1.5"
    >
        <Handle type="target" :position="Position.Left" id="left" />
        <Handle type="target" :position="Position.Right" id="right" />
        <Handle
            type="source"
            :position="Position.Bottom"
            :class="{ invisible: !bottomHandleConnected }"
        />

        <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="3"
            stroke="currentColor"
            class="size-6"
        >
            <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M13.19 8.688a4.5 4.5 0 0 1 1.242 7.244l-4.5 4.5a4.5 4.5 0 0 1-6.364-6.364l1.757-1.757m13.35-.622 1.757-1.757a4.5 4.5 0 0 0-6.364-6.364l-4.5 4.5a4.5 4.5 0 0 0 1.242 7.244"
            />
        </svg>
    </div>
</template>

<script setup lang="ts">
    import { computed, ComputedRef } from "vue";
    import { Position, Handle, useVueFlow } from "@vue-flow/core";
    import type { NodeProps } from "@vue-flow/core";

    // Props
    const props = defineProps<NodeProps>();

    // Composables
    const { edges } = useVueFlow();

    // Computed properties
    const bottomHandleConnected: ComputedRef<boolean> = computed(() => {
        return edges.value.some((edge) => edge.source == props.id);
    });
</script>

<style lang="scss" scoped>
    .node-root {
        width: 32px;
        height: 32px;
    }
</style>
