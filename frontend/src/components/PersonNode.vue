<template>
    <div class="vue-flow__node-default node-root" :class="[ data.sex === 'male' ? '!bg-blue-200' : '!bg-red-200' ]">
        <Handle type="target" :position="Position.Top" id="top" :class="{'invisible' : !connectedHandles.top}"/>
        <Handle type="source" :position="Position.Left" id="left" :class="{'invisible' : !connectedHandles.left}"/>
        <Handle type="source" :position="Position.Right" id="right" :class="{'invisible' : !connectedHandles.right}"/>

        <div class="aspect-square bg-white border-b">
            <img :src="profilePicture" alt="profile picture"/>
        </div>

        <div class="bg-transparent">{{ data.firstname }} {{ data.lastname }}</div>
    </div>
</template>

<script setup lang="ts">
import {Position, Handle, useVueFlow} from '@vue-flow/core'
import type { NodeProps } from '@vue-flow/core'
import {computed, ComputedRef} from "vue";
import PersonPlaceholder from "../assets/images/person_placeholder.svg";

// Props
const props = defineProps<NodeProps>();

// Composables
const { edges } = useVueFlow();

// Computed properties
const connectedHandles = computed(() => {
    let ownEdges = edges.value.filter(edge => [edge.source, edge.target].includes(props.id));

    return {
        top: ownEdges.some(edge => edge.target == props.id) ?? false,
        left: ownEdges.some(edge => edge.sourceHandle == 'left') ?? false,
        right: ownEdges.some(edge => edge.sourceHandle == 'right') ?? false,
    }
});

const profilePicture: ComputedRef<string> = computed(() => {
    return PersonPlaceholder;
});

</script>

<style lang="scss" scoped>
    .node-root {
        padding: 0 !important;
        width: 100px;
        height: 150px;
    }
</style>