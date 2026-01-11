<template>
    <div class="h-full max-h-full p-4 bg-gray-200 overflow-auto">
        <h2 class="font-bold text-xl">Editor</h2>

        <hr class="mb-4" />

        <template v-if="selectedNode">
            <div class="mb-2">
                <label for="firstname" class="font-semibold">First name</label>
                <input
                    v-model="selectedNode.data.firstname"
                    name="firstname"
                    type="text"
                    class="w-full bg-white border border-solid border-gray-400
                        px-1"
                />
            </div>

            <div class="mb-2">
                <label for="lastname" class="font-semibold">Last name</label>
                <input
                    v-model="selectedNode.data.lastname"
                    name="lastname"
                    type="text"
                    class="w-full bg-white border border-solid border-gray-400
                        px-1"
                />
            </div>

            <div class="mb-2">
                <label for="sex" class="font-semibold">Sex</label>
                <select
                    v-model="selectedNode.data.sex"
                    name="sex"
                    class="w-full bg-white border border-solid border-gray-400 px-1"
                >
                    <option key="male" value="male">Male</option>
                    <option key="female" value="female">Female</option>
                </select>
            </div>

            <div class="mb-2">
                <label for="birthdate" class="font-semibold">Date of birth</label>
                <!-- :unmask="true" -->
                <IMaskComponent
                    v-model="selectedNode.data.birthdate"
                    mask="0000-00-00"
                    name="birthdate"
                    type="text"
                    placeholder="1950-12-01"
                    class="w-full bg-white border border-solid border-gray-400
                        px-1"
                />
            </div>

            <div class="mb-2">
                <label for="birthplace" class="font-semibold">Location of birth</label>
                <input
                    v-model="selectedNode.data.birthplace"
                    name="birthplace"
                    type="text"
                    class="w-full bg-white border border-solid border-gray-400
                        px-1" />
            </div>
        </template>

        <template v-else> Select a person to edit </template>
    </div>
</template>

<script lang="ts" setup>
    import { Node } from "@vue-flow/core";
    import { onMounted, ref, watch } from "vue";
    import { IMaskComponent } from "vue-imask";

    // Props
    const props = defineProps<{
        selectedNode: Node | null;
    }>();
</script>
