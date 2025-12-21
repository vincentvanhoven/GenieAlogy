<template>
    <div class="h-full max-h-full p-4 bg-gray-200 overflow-auto">
        <h2 class="font-bold text-xl">Editor</h2>

        <hr class="mb-4" />

        <template v-if="selectedNode">
            <div class="mb-4">ID #{{ selectedNode?.data.id }}</div>

            <div class="mb-4">
                <label for="firstname" class="font-semibold">First name</label>
                <input
                    v-model="selectedNode.data.firstname"
                    name="firstname"
                    type="text"
                    class="w-full bg-white border border-solid border-gray-400"
                />
            </div>

            <div class="mb-4">
                <label for="lastname" class="font-semibold">Last name</label>
                <input
                    v-model="selectedNode.data.lastname"
                    name="lastname"
                    type="text"
                    class="w-full bg-white border border-solid border-gray-400"
                />
            </div>
        </template>

        <template v-else> Select a person to edit </template>
    </div>
</template>

<script lang="ts" setup>
    import { Node } from "@vue-flow/core";

    // Props
    const props = defineProps<{
        selectedNode: Node | null;
    }>();
</script>
