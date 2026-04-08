<template>
    <Dialog
        v-model:visible="saveFileStore.isAddingFamily"
        modal
        header="Add partnership"
        :style="{ width: '25rem' }"
    >
        <div class="">
            <label for="person_a">Person A</label>
            <Select
                size="small"
                :options="saveFileStore.people"
                :default-value="saveFileStore.selectedPerson"
                :optionLabel="saveFileStore.getPersonDisplayName"
                class="w-full mb-2"
                :disabled="true"
            />

            <label for="person_a">Person B</label>
            <Select
                size="small"
                v-model="personB"
                :options="potentialPeopleB"
                :optionLabel="saveFileStore.getPersonDisplayName"
                class="w-full mb-4"
                :disabled="false"
            />
        </div>

        <div class="flex justify-between gap-2">
            <Button
                type="button"
                label="Cancel"
                severity="secondary"
                @click="cancel"
            ></Button>
            <Button type="button" label="Add" @click="confirm" :disabled="!personB"></Button>
        </div>
    </Dialog>
</template>

<script lang="ts" setup>
    import { Ref, ref, ComputedRef, computed } from "vue";
    import { Button, Dialog, FloatLabel, Select } from "primevue";
    import { useSaveFileStore } from "../../stores/saveFileStore";
    import { models } from "../../../wailsjs/go/models";
    import Person = models.Person;
    import Family = models.Family;

    // Data
    const saveFileStore = useSaveFileStore();
    const personB: Ref<Person | null> = ref(null);

    // Computed
    const potentialPeopleB: ComputedRef<Person[]> = computed(() => {
        let personBSex = saveFileStore.selectedPerson!.sex === 'male' ? 'female' : 'male';
        return saveFileStore.people!.filter((person) => person.sex === personBSex) ?? [];
    });

    // Methods
    function cancel() {
        saveFileStore.disableAddFamilyMode();
    }

    function confirm() {
        if (!personB.value) {
            return;
        }

        let personA = saveFileStore.selectedPerson!;

        // Compute the midpoint between both People for the Partnership node to be placed. Note that the -16 (x) and
        // +32 (y) take the dimensions of the Partnership node into account.
        const cx = Math.round((personA.position_x + personB.value.position_x - 32 + saveFileStore.personNodeDimensions.x) / 2 / 16) * 16;
        const cy = Math.round((personA.position_y + personB.value.position_y + 32) / 2 / 16) * 16;

        let family: Family = {
            id: undefined,
            male_id: personA.sex === 'male' ? personA.id : personB.value.id,
            female_id: personA.sex === 'female' ? personA.id : personB.value.id,
            position_x: cx,
            position_y: cy,
        };

        saveFileStore.disableAddFamilyMode(family);
    }
</script>
