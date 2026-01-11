<template>
    <div
        class="vue-flow__node-default node-root"
        :class="[data.sex === 'male' ? '!bg-blue-200' : '!bg-red-200']"
    >
        <Handle
            type="target"
            :position="Position.Top"
            :class="{ invisible: !connectedHandles.top }"
        />

        <Handle
            v-if="data.sex === 'female'"
            type="source"
            :position="Position.Left"
            :class="{ invisible: !connectedHandles.left }"
        />

        <Handle
            v-else
            type="source"
            :position="Position.Right"
            id="right"
            :class="{ invisible: !connectedHandles.right }"
        />

        <div
            class="h-full flex justify-center items-center text-sm font-semibold
                bg-transparent"
        >
            {{ data.firstname }} {{ data.lastname }} ({{ data.id }})
        </div>
    </div>
</template>

<script setup lang="ts">
    import { Position, Handle, useVueFlow } from "@vue-flow/core";
    import type { NodeProps } from "@vue-flow/core";
    import { computed, ComputedRef } from "vue";
    import PersonPlaceholder from "../assets/images/person_placeholder.svg";

    // Props
    const props = defineProps<NodeProps>();

    // Composables
    const { edges } = useVueFlow();

    // Computed properties
    const connectedHandles = computed(() => {
        let ownEdges = edges.value.filter((edge) =>
            [edge.source, edge.target].includes(props.id),
        );

        let hasTargetEdge = ownEdges.some((edge) => edge.target == props.id);
        let hasSourceEdge = ownEdges.some((edge) => edge.source == props.id);

        return {
            top: hasTargetEdge ?? false,
            left: (hasSourceEdge && props.data.sex === "female") ?? false,
            right: (hasSourceEdge && props.data.sex === "male") ?? false,
        };
    });

    const profilePicture: ComputedRef<string> = computed(() => {
        return PersonPlaceholder;
    });
</script>

<style lang="scss" scoped>
    .node-root {
        padding: 0 !important;
        width: 256px;
        height: 64px;
    }
</style>
