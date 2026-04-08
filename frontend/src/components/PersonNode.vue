<template>
    <div
        class="vue-flow__node-default node-root"
        :class="[selfPerson.sex === 'male' ? '!bg-blue-200' : '!bg-red-200']"
    >
        <Handle
            type="target"
            :position="targetHandlePosition"
            :class="{ invisible: !connectedHandles.top }"
        />
        <Handle
            v-if="selfPerson.sex === 'female'"
            type="source"
            id="right"
            :position="Position.Right"
            :class="{ invisible: !connectedHandles.right }"
        />
        <Handle
            v-else
            type="source"
            id="left"
            :position="Position.Left"
            :class="{ invisible: !connectedHandles.left }"
        />
        <div
            class="h-full flex justify-center items-center text-sm font-semibold
                bg-transparent"
        >
            {{ saveFileStore.getPersonDisplayName(selfPerson) }}
        </div>
    </div>
</template>

<script setup lang="ts">
    import { Handle, useVueFlow, Position } from "@vue-flow/core";
    import type { NodeProps } from "@vue-flow/core";
    import { computed, ComputedRef } from "vue";
    import { useSaveFileStore } from "../stores/saveFileStore";
    import { models } from "../../wailsjs/go/models";
    import Person = models.Person;

    // Props
    const props = defineProps<NodeProps>();

    // Composables
    const { edges } = useVueFlow();
    const saveFileStore = useSaveFileStore();

    // Computed properties
    const selfPerson: ComputedRef<Person> = computed(() => {
        return saveFileStore.getPersonFromNode(props.id) as Person;
    });

    const connectedHandles = computed(() => {
        let ownEdges = edges.value.filter((edge) =>
            [edge.source, edge.target].includes(props.id),
        );

        let hasSourceEdge = ownEdges.some((edge) => edge.source == props.id);

        return {
            top: !!selfPerson.value.family_id,
            // TODO: Refactor this to be more dynamic
            left: (hasSourceEdge && selfPerson.value.sex === "male") ?? false,
            right:
                (hasSourceEdge && selfPerson.value.sex === "female") ?? false,
        };
    });

    const targetHandlePosition = computed(() => {
        let position = selfPerson.value?.parent_arrow_position ?? "top";

        if (position == "top") {
            return Position.Top;
        } else if (position == "left") {
            return Position.Left;
        } else if (position == "right") {
            return Position.Right;
        }
    });
</script>

<style lang="scss" scoped>
    .node-root {
        padding: 0 !important;
        width: 256px;
        height: 64px;
    }
</style>
