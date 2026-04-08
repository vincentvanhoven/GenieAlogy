<template>
    <div
        class="w-full h-full max-h-screen border-l border-l-gray-300 shadow
            text-sm"
    >
        <Form
            class="h-full max-h-full flex flex-col"
            v-if="saveFileStore.selectedPerson"
            :key="saveFileStore.selectedPerson.id"
            :initial-values="saveFileStore.selectedPerson"
            @submit="onSubmit"
        >
            <div
                class="px-4 py-2 border-b border-b-gray-200 flex
                    justify-between"
            >
                <h2 class="font-bold text-xl">Editor</h2>

                <div>
                    <Button
                        label="Remove"
                        class="leading-4 mr-2"
                        :class="{
                            'cursor-not-allowed':
                                saveFileStore.isEditingPerson ||
                                saveFileStore.isDeletingPerson,
                        }"
                        severity="danger"
                        size="small"
                        :disabled="
                            saveFileStore.isEditingPerson ||
                            saveFileStore.isDeletingPerson
                        "
                        @click="saveFileStore.enableDeletePersonMode"
                    />

                    <Button
                        label="Edit"
                        class="leading-4"
                        :class="{
                            'cursor-not-allowed': saveFileStore.isEditingPerson,
                        }"
                        severity="contrast"
                        size="small"
                        :disabled="saveFileStore.isEditingPerson"
                        @click="saveFileStore.enableEditMode"
                    />
                </div>
            </div>

            <div class="overflow-auto flex-1 px-4 pb-4">
                <Divider align="left" type="solid" size="small">
                    <b>Basic information</b>
                </Divider>

                <FloatLabel variant="on" class="flex-1 mb-4">
                    <label for="firstname">First Name</label>
                    <InputText
                        size="small"
                        id="firstname"
                        name="firstname"
                        class="w-full"
                        :disabled="!saveFileStore.isEditingPerson"
                    />
                </FloatLabel>

                <FloatLabel variant="on" class="flex-1 mb-4">
                    <label for="lastname">Last Name</label>
                    <InputText
                        size="small"
                        id="lastname"
                        name="lastname"
                        class="w-full"
                        :disabled="!saveFileStore.isEditingPerson"
                    />
                </FloatLabel>

                <FloatLabel variant="on">
                    <Select
                        size="small"
                        id="sex"
                        name="sex"
                        :options="[
                            { label: 'Male', value: 'male' },
                            { label: 'Female', value: 'female' },
                        ]"
                        optionLabel="label"
                        optionValue="value"
                        class="w-full"
                        :disabled="!saveFileStore.isEditingPerson"
                    />
                    <label for="sex">Sex</label>
                </FloatLabel>

                <Divider align="left" type="solid">
                    <b>Birth</b>
                </Divider>

                <FloatLabel variant="on" class="mb-4">
                    <label for="birthdate">Date of Birth</label>
                    <DatePicker
                        size="small"
                        id="birthdate"
                        name="birthdate"
                        dateFormat="yy-mm-dd"
                        class="w-full"
                        :disabled="!saveFileStore.isEditingPerson"
                        :model-value="birthdate"
                    />
                </FloatLabel>

                <FloatLabel variant="on">
                    <label for="birthplace">Location of Birth</label>
                    <InputText
                        size="small"
                        id="birthplace"
                        name="birthplace"
                        class="w-full"
                        :disabled="!saveFileStore.isEditingPerson"
                    />
                </FloatLabel>

                <Divider align="left" type="solid">
                    <b>Death</b>
                </Divider>

                <FloatLabel variant="on" class="mb-4">
                    <label for="deathdate">Date of Death</label>
                    <DatePicker
                        size="small"
                        id="deathdate"
                        name="deathdate"
                        dateFormat="yy-mm-dd"
                        class="w-full"
                        :disabled="!saveFileStore.isEditingPerson"
                        :model-value="deathdate"
                    />
                </FloatLabel>

                <FloatLabel variant="on" class="mb-4">
                    <label for="deathplace">Location of Death</label>
                    <InputText
                        size="small"
                        id="deathplace"
                        name="deathplace"
                        class="w-full"
                        :disabled="!saveFileStore.isEditingPerson"
                    />
                </FloatLabel>

                <Divider align="left" type="solid">
                    <b>Family</b>
                </Divider>

                <FloatLabel variant="on">
                    <Select
                        size="small"
                        id="family_id"
                        name="family_id"
                        :options="familyOptions"
                        optionLabel="name"
                        optionValue="id"
                        class="w-full mb-4"
                        :disabled="!saveFileStore.isEditingPerson"
                    />
                    <label for="family_id">Parents</label>
                </FloatLabel>

                <FloatLabel variant="on">
                    <Select
                        size="small"
                        id="parent_arrow_position"
                        name="parent_arrow_position"
                        :options="[
                            { name: 'Top', value: 'top' },
                            { name: 'Left', value: 'left' },
                            { name: 'Right', value: 'right' },
                        ]"
                        optionLabel="label"
                        optionValue="value"
                        class="w-full"
                        :disabled="!saveFileStore.isEditingPerson"
                    />
                    <label for="parent_arrow_position">Arrow position</label>
                </FloatLabel>

                <Divider align="left" type="solid" class="mb-0">
                    <b>Partnerships</b>
                </Divider>

                <DataTable :value="partnerships" dataKey="id">
                    <template #empty> No partnerships </template>

                    <Column field="name" header="Name"></Column>

                    <Column class="w-8">
                        <template #header>
                            <Button
                                :class="{
                                    'cursor-not-allowed':
                                        !saveFileStore.isEditingPerson,
                                }"
                                :disabled="!saveFileStore.isEditingPerson"
                                severity="contrast"
                                size="small"
                                icon="pi pi-plus"
                                aria-label="Add"
                                class="p-1 w-6 h-6"
                                @click="saveFileStore.enableAddFamilyMode"
                            />
                        </template>

                        <template #body="{ data }">
                            <Button
                                :class="{
                                    'cursor-not-allowed':
                                        !saveFileStore.isEditingPerson,
                                }"
                                :disabled="!saveFileStore.isEditingPerson"
                                severity="danger"
                                size="small"
                                icon="pi pi-trash"
                                aria-label="Remove"
                                class="p-1 w-6 h-6 mr-1"
                                @click="removeFamily(data.id)"
                            />

                            <!--
                            <Button
                                :class="{'cursor-not-allowed' : !saveFileStore.isEditingPerson }"
                                :disabled="!saveFileStore.isEditingPerson"
                                severity="contrast"
                                size="small"
                                icon="pi pi-pencil"
                                aria-label="Edit"
                                class="p-1 w-6 h-6"
                            />
                            -->
                        </template>
                    </Column>
                </DataTable>
            </div>

            <div
                v-if="saveFileStore.isEditingPerson"
                class="px-4 py-2 border-t border-t-gray-200 flex
                    justify-between"
            >
                <Button
                    label="Cancel"
                    class="leading-4"
                    severity="danger"
                    @click="cancelEdit"
                />
                <Button type="submit" label="Save" class="leading-4" />
            </div>
        </Form>

        <template v-else>
            <div class="px-4 py-2">Select a person to edit</div>
        </template>
    </div>
</template>

<script lang="ts" setup>
    import { Form, FormSubmitEvent } from "@primevue/forms";
    import {
        Button,
        DatePicker,
        Divider,
        FloatLabel,
        InputText,
        Select,
        DataTable,
        Column,
    } from "primevue";
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
            name: saveFileStore.getFamilyDisplayName(family),
        }));
    });

    // Methods
    function onSubmit(formSubmitEvent: FormSubmitEvent) {
        let anyFieldsChanged = Object.values(formSubmitEvent.states).filter(
            (value) => value.dirty,
        );
        let payload: Person | null = anyFieldsChanged
            ? (formSubmitEvent.values as Person)
            : null;

        if (payload) {
            payload.birthdate = payload.birthdate
                ? dayjs(payload.birthdate).format("YYYY-MM-DD")
                : undefined;
        }

        saveFileStore.disableEditMode(payload);
    }

    function cancelEdit() {
        saveFileStore.disableEditMode();
    }

    function removeFamily(familyId: number) {
        let family = saveFileStore.families.find(
            (family) => family.id === familyId,
        )!;
        saveFileStore.enableDeleteFamilyMode(family);
    }

    // Computed properties
    const partnerships = computed(() => {
        return saveFileStore.families
            .filter((family) => {
                return (
                    family.male_id === saveFileStore.selectedPerson?.id ||
                    family.female_id === saveFileStore.selectedPerson?.id
                );
            })
            .map((family) => {
                let otherPersonId =
                    family.male_id === saveFileStore.selectedPerson?.id
                        ? family.female_id
                        : family.male_id;
                let otherPerson = saveFileStore.getPersonFromId(
                    otherPersonId!,
                )!;

                return {
                    id: family.id,
                    name: saveFileStore.getPersonDisplayName(otherPerson),
                    type: "Unknown",
                    start_date: "1900-01-01",
                    end_date: "1900-01-01",
                };
            });
    });

    const birthdate = computed(
        () => new Date(saveFileStore.selectedPerson?.birthdate ?? ""),
    );
    const deathdate = computed(
        () => new Date(saveFileStore.selectedPerson?.deathdate ?? ""),
    );
</script>
