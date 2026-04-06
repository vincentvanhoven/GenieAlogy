<template>
    <div
        class="vue-flow__node-default node-root flex justify-center items-center
            !rounded-full !border !p-1.5"
    >
        <Handle type="target" :position="Position.Left" id="left" />
        <Handle type="target" :position="Position.Right" id="right" />
        <Handle
            type="source"
            :position="Position.Bottom"
            :class="{ invisible: !bottomHandleConnected }"
        />

        <img src="../assets/images/link.svg" class="w-5 h-5" />
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
