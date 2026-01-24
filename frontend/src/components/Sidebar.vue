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
                    class="w-full bg-white border border-solid border-gray-400
                        px-1"
                >
                    <option key="male" value="male">Male</option>
                    <option key="female" value="female">Female</option>
                </select>
            </div>

            <div class="mb-2">
                <label for="birthdate" class="font-semibold"
                >Date of birth</label
                >

                <IMaskComponent
                    v-model="birthdate"
                    mask="0000-00-00"
                    name="birthdate"
                    type="text"
                    placeholder="1950-12-01"
                    class="w-full bg-white border border-solid border-gray-400
                        px-1"
                />
            </div>

            <div class="mb-2">
                <label for="birthplace" class="font-semibold"
                >Location of birth</label
                >
                <input
                    v-model="selectedNode.data.birthplace"
                    name="birthplace"
                    type="text"
                    class="w-full bg-white border border-solid border-gray-400
                        px-1"
                />
            </div>

            <div class="mb-2">
                <label for="family_id" class="font-semibold">Parents</label>
                <select
                    v-model="selectedNode.data.family_id"
                    name="family_id"
                    class="w-full bg-white border border-solid border-gray-400
                        px-1"
                >
                    <option key="null" :value="null"></option>
                    <option
                        v-for="family in families"
                        :key="family.id"
                        :value="family.id"
                    >
                        {{ formatFamily(family) }}
                    </option>
                </select>
            </div>

            <div class="mb-2">
                <table>
                    <thead>
                    <tr>
                        <th class="w-full text-left">Partnerships</th>
                        <th class="pl-2">
                            <button
                                @click="startAddingFamily"
                                type="button"
                                class="rounded-full border border-gray-500 p-0.5 cursor-pointer
                                    hover:border-gray-400 hover:bg-gray-200 hover:shadow
                                    active:bg-gray-300 w-5 h-5 flex items-center justify-center"
                            >
                                <img src="../assets/images/plus.svg" class="w-4 h-4" />
                            </button>
                        </th>
                    </tr>
                    </thead>

                    <tbody>
                    <tr v-for="family in partnerships">
                        <td>
                            {{ formatFamily(family) }}
                        </td>
                        <td class="pl-2">
                            <button
                                @click="removeFamily(family)"
                                type="button"
                                class="rounded-full border border-gray-500 p-0.5 cursor-pointer
                                    hover:border-gray-400 hover:bg-gray-200 hover:shadow
                                    active:bg-gray-300 w-5 h-5 flex items-center justify-center"
                            >
                                <img src="../assets/images/trash.svg" class="w-4 h-4" />
                            </button>
                        </td>
                    </tr>
                    <tr v-if="addingFamily">
                        <td>
                            <select
                                v-model="newFamilyPartnerId"
                                name="new_family_partner_id"
                                class="w-full bg-white border border-solid border-gray-400 px-1"
                            >
                                <option key="null" :value="null"></option>
                                <option
                                    v-for="person in people.filter((p: Person) => p.id !== selectedNode?.data.id)"
                                    :key="person.id"
                                    :value="person.id"
                                >
                                    {{ person.firstname }} ({{ person.id}})
                                </option>
                            </select>
                        </td>
                        <td class="pl-2">
                            <button
                                @click="finishAddingFamily"
                                type="button"
                                class="rounded-full border border-gray-500 p-0.5 cursor-pointer
                                    hover:border-gray-400 hover:bg-gray-200 hover:shadow
                                    active:bg-gray-300 w-5 h-5 flex items-center justify-center"
                            >
                                <img src="../assets/images/check.svg" class="w-4 h-4" />
                            </button>
                        </td>
                    </tr>
                    </tbody>
                </table>
            </div>
        </template>

        <template v-else> Select a person to edit</template>
    </div>
</template>

<script lang="ts" setup>
    import { Node } from "@vue-flow/core";
    import { models } from "../../wailsjs/go/models";
    import { computed, ref, Ref } from "vue";
    import { IMaskComponent } from "vue-imask";
    import Family = models.Family;
    import Person = models.Person;

    // Props
    const props = defineProps<{
        selectedNode: Node | null;
        people: Person[];
        families: Family[];
    }>();

    // Data
    const addingFamily: Ref<boolean> = ref(false);
    const newFamilyPartnerId: Ref<number|null> = ref(null);

    // Events
    const emit = defineEmits<{
        (e: 'add-family', family: Family): void;
        (e: 'remove-family', family: Family): void;
    }>()

    // Computed properties
    const birthdate = computed({
        get() {
            return props.selectedNode?.data.birthdate || "";
        },
        set(value) {
            if (props.selectedNode) {
                props.selectedNode.data.birthdate = value || null;
            }
        }
    });

    const partnerships = computed(() => {
        return props.families.filter((family) => {
            return family.person_1_id === props.selectedNode?.data.id ||
                family.person_2_id === props.selectedNode?.data.id;
        });
    });

    // Methods
    function formatFamily(family: Family): string {
        let parents = props.people.filter((person) => {
            return (
                family.person_1_id === person.id ||
                family.person_2_id === person.id
            );
        });

        return `${parents[0].firstname} (${parents[0].id}) - ${parents[1].firstname} (${parents[1].id})`;
    }

    function startAddingFamily() {
        addingFamily.value = true;
    }

    function finishAddingFamily() {
        if (!newFamilyPartnerId.value) {
            return;
        }

        let selfPerson: Person = props.selectedNode?.data;
        let newFamily: Family = {
            person_1_id: selfPerson.sex == "male" ? selfPerson.id : newFamilyPartnerId.value,
            person_2_id: selfPerson.sex == "male" ? newFamilyPartnerId.value : selfPerson.id,
            position_x: 0,
            position_y: 0,
        }

        emit('add-family', newFamily);
        addingFamily.value = false;
        newFamilyPartnerId.value = null;
    }

    function removeFamily(family: Family) {
        emit('remove-family', family);
    }
</script>
