<template>
    <div class="w-full h-full max-h-screen border-l border-l-gray-300 shadow">
        <Form
            v-if="saveFileStore.selectedPerson"
            :key="saveFileStore.selectedPerson.id"
            :initial-values="saveFileStore.selectedPerson"
            @submit="onSubmit"
            class="h-full max-h-full flex flex-col"
        >
            <div class="px-4 py-2 border-b border-b-gray-200 flex justify-between">
                <h2 class="font-bold text-xl">Editor</h2>

                <Button
                    label="Edit"
                    class="leading-4"
                    :class="{'cursor-not-allowed' : saveFileStore.isEditingPerson }"
                    severity="contrast"
                    size="small"
                    :disabled="saveFileStore.isEditingPerson"
                    @click="saveFileStore.enableEditMode"
                />
            </div>

            <div class="overflow-auto flex-1 px-4 pb-4">
                <Divider align="left" type="solid"><b>Basic information</b></Divider>

                <FloatLabel variant="on" class="flex-1 mb-4">
                    <label for="firstname">First Name</label>
                    <InputText
                        id="firstname"
                        name="firstname"
                        class="w-full"
                        :disabled="!saveFileStore.isEditingPerson"
                    />
                </FloatLabel>

                <FloatLabel variant="on" class="flex-1 mb-4">
                    <label for="lastname">Last Name</label>
                    <InputText
                        id="lastname"
                        name="lastname"
                        class="w-full"
                        :disabled="!saveFileStore.isEditingPerson"
                    />
                </FloatLabel>

                <FloatLabel variant="on">
                    <Select
                        id="sex"
                        name="sex"
                        :options="[{ label: 'Male', value: 'male' }, { label: 'Female', value: 'female' }]"
                        optionLabel="label"
                        optionValue="value"
                        class="w-full"
                        :disabled="!saveFileStore.isEditingPerson"
                    />
                    <label for="sex">Sex</label>
                </FloatLabel>

                <Divider align="left" type="solid"><b>Birth</b></Divider>

                <FloatLabel variant="on" class="mb-4">
                    <label for="birthdate">Date of Birth</label>
                    <DatePicker
                        id="birthdate"
                        name="birthdate"
                        dateFormat="yy-mm-dd"
                        class="w-full"
                        :disabled="!saveFileStore.isEditingPerson"
                    />
                </FloatLabel>

                <FloatLabel variant="on">
                    <label for="birthplace">Location of Birth</label>
                    <InputText
                        id="birthplace"
                        name="birthplace"
                        class="w-full"
                        :disabled="!saveFileStore.isEditingPerson"
                    />
                </FloatLabel>

                <!--
                                <Divider align="left" type="solid"><b>Death</b></Divider>

                                <FloatLabel variant="on" class="mb-4">
                                    <label for="birthdate">Date of Birth</label>
                                    <DatePicker id="birthdate" name="birthdate" dateFormat="yy-mm-dd" class="w-full" />
                                </FloatLabel>

                                <FloatLabel variant="on" class="mb-4">
                                    <label for="birthplace">Location of Birth</label>
                                    <InputText id="birthplace" name="birthplace" class="w-full" />
                                </FloatLabel>
                -->

                <Divider align="left" type="solid"><b>Family</b></Divider>

                <FloatLabel variant="on">
                    <Select
                        id="family_id"
                        name="family_id"
                        :options="familyOptions"
                        optionLabel="name"
                        optionValue="id"
                        class="w-full"
                        :disabled="!saveFileStore.isEditingPerson"
                    />
                    <label for="family_id">Parents</label>
                </FloatLabel>
            </div>

            <div v-if="saveFileStore.isEditingPerson" class="px-4 py-2 border-t border-t-gray-200 flex justify-between">
                <Button label="Cancel" class="leading-4" severity="danger" @click="cancelEdit" />
                <Button type="submit" label="Save" class="leading-4"/>
            </div>
        </Form>

        <!--            <div class="mb-2">-->
        <!--                <table>-->
        <!--                    <thead>-->
        <!--                        <tr>-->
        <!--                            <th class="w-full text-left">Partnerships</th>-->
        <!--                            <th class="pl-2">-->
        <!--                                <button-->
        <!--                                    @click="startAddingFamily"-->
        <!--                                    type="button"-->
        <!--                                    class="rounded-full border border-gray-500-->
        <!--                                        p-0.5 cursor-pointer-->
        <!--                                        hover:border-gray-400 hover:bg-gray-200-->
        <!--                                        hover:shadow active:bg-gray-300 w-5 h-5-->
        <!--                                        flex items-center justify-center"-->
        <!--                                >-->
        <!--                                    <img-->
        <!--                                        src="../assets/images/plus.svg"-->
        <!--                                        class="w-4 h-4"-->
        <!--                                    />-->
        <!--                                </button>-->
        <!--                            </th>-->
        <!--                        </tr>-->
        <!--                    </thead>-->

        <!--                    <tbody>-->
        <!--                        <tr v-for="family in partnerships">-->
        <!--                            <td>-->
        <!--                                {{ formatFamily(family) }}-->
        <!--                            </td>-->
        <!--                            <td class="pl-2">-->
        <!--                                <button-->
        <!--                                    @click="removeFamily(family)"-->
        <!--                                    type="button"-->
        <!--                                    class="rounded-full border border-gray-500-->
        <!--                                        p-0.5 cursor-pointer-->
        <!--                                        hover:border-gray-400 hover:bg-gray-200-->
        <!--                                        hover:shadow active:bg-gray-300 w-5 h-5-->
        <!--                                        flex items-center justify-center"-->
        <!--                                >-->
        <!--                                    <img-->
        <!--                                        src="../assets/images/trash.svg"-->
        <!--                                        class="w-4 h-4"-->
        <!--                                    />-->
        <!--                                </button>-->
        <!--                            </td>-->
        <!--                        </tr>-->
        <!--                        <tr v-if="addingFamily">-->
        <!--                            <td>-->
        <!--                                <select-->
        <!--                                    v-model="newFamilyPartnerId"-->
        <!--                                    name="new_family_partner_id"-->
        <!--                                    class="w-full bg-white border border-solid-->
        <!--                                        border-gray-400 px-1"-->
        <!--                                >-->
        <!--                                    <option key="null" :value="null"></option>-->
        <!--                                    <option-->
        <!--                                        v-for="person in people.filter(-->
        <!--                                            (p: Person) =>-->
        <!--                                                p.id !== selectedNode?.data.id,-->
        <!--                                        )"-->
        <!--                                        :key="person.id"-->
        <!--                                        :value="person.id"-->
        <!--                                    >-->
        <!--                                        {{ person.firstname }} ({{ person.id }})-->
        <!--                                    </option>-->
        <!--                                </select>-->
        <!--                            </td>-->
        <!--                            <td class="pl-2">-->
        <!--                                <button-->
        <!--                                    @click="finishAddingFamily"-->
        <!--                                    type="button"-->
        <!--                                    class="rounded-full border border-gray-500-->
        <!--                                        p-0.5 cursor-pointer-->
        <!--                                        hover:border-gray-400 hover:bg-gray-200-->
        <!--                                        hover:shadow active:bg-gray-300 w-5 h-5-->
        <!--                                        flex items-center justify-center"-->
        <!--                                >-->
        <!--                                    <img-->
        <!--                                        src="../assets/images/check.svg"-->
        <!--                                        class="w-4 h-4"-->
        <!--                                    />-->
        <!--                                </button>-->
        <!--                            </td>-->
        <!--                        </tr>-->
        <!--                    </tbody>-->
        <!--                </table>-->
        <!--            </div>-->

        <template v-else>
            <div class="px-4 py-2">Select a person to edit</div>
        </template>
    </div>
</template>

<script lang="ts" setup>
    import { Form, FormSubmitEvent } from "@primevue/forms";
    import { Button, DatePicker, Divider, FloatLabel, InputText, Select } from "primevue";
    import { useSaveFileStore } from "../stores/saveFileStore";
    import { computed } from "vue";
    import { models } from "../../wailsjs/go/models";
    import Person = models.Person;
    import dayjs from "dayjs";

    // Data
    const saveFileStore = useSaveFileStore();

    // Computed properties
    const familyOptions = computed(() => {
        return saveFileStore.families.map((family) => ({
            id: family.id,
            name: saveFileStore.getFamilyDisplayName(family)
        }));
    });

    // Methods
    function onSubmit(formSubmitEvent: FormSubmitEvent) {
        let anyFieldsChanged = Object.values(formSubmitEvent.states).filter(value => value.dirty);
        let payload: Person|null = anyFieldsChanged ? formSubmitEvent.values as Person : null

        if (payload) {
            payload.birthdate = payload.birthdate ? dayjs(payload.birthdate).format("YYYY-MM-DD") : undefined;
        }

        saveFileStore.disableEditMode(payload);
    }

    function cancelEdit() {
        saveFileStore.disableEditMode();
    }

    // Computed properties
    // const birthdate = computed({
    //     get() {
    //         return saveFileStore.selectedPerson?.birthdate || "";
    //     },
    //     set(value) {
    //         if (saveFileStore.selectedPerson) {
    //             saveFileStore.selectedPerson.birthdate = value || null;
    //         }
    //     },
    // });

    // const partnerships = computed(() => {
    //     return props.families.filter((family) => {
    //         return (
    //             family.male_id === props.selectedNode?.data.id ||
    //             family.female_id === props.selectedNode?.data.id
    //         );
    //     });
    // });

    // Methods
    // function startAddingFamily() {
    //     addingFamily.value = true;
    // }

    // function finishAddingFamily() {
    //     if (!newFamilyPartnerId.value) {
    //         return;
    //     }
    //
    //     let selfPerson: Person = props.selectedNode?.data;
    //     let newFamily: Family = {
    //         male_id:
    //             selfPerson.sex == "male"
    //                 ? selfPerson.id
    //                 : newFamilyPartnerId.value,
    //         female_id:
    //             selfPerson.sex == "male"
    //                 ? newFamilyPartnerId.value
    //                 : selfPerson.id,
    //         position_x: 0,
    //         position_y: 0,
    //     };
    //
    //     emit("add-family", newFamily);
    //     addingFamily.value = false;
    //     newFamilyPartnerId.value = null;
    // }
    //
    // function removeFamily(family: Family) {
    //     emit("remove-family", family);
    // }
</script>
