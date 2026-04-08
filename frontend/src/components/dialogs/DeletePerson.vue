<template>
    <Dialog v-model:visible="saveFileStore.isDeletingPerson" modal
        header="Confirm removal"
        :style="{ width: '25rem' }"
    >
        <span class="block mb-8">
            This person will be replaced with an anonymized card if they have any descendants.
        </span>

        <div class="flex justify-between gap-2">
            <Button type="button" label="Cancel" severity="secondary" @click="cancel"></Button>
            <Button type="button" label="Remove" severity="danger" @click="confirm"></Button>
        </div>
    </Dialog>
</template>

<script lang="ts" setup>
    import { Button, Dialog } from "primevue";
    import { useSaveFileStore } from "../../stores/saveFileStore";

    // Data
    const saveFileStore = useSaveFileStore();

    // Methods
    function cancel() {
        saveFileStore.disableDeletePersonMode();
    }

    function confirm() {
        saveFileStore.disableDeletePersonMode(true);
    }
</script>