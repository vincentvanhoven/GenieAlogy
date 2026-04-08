<template>
    <Dialog
        v-model:visible="saveFileStore.isDeletingFamily"
        modal
        header="Confirm removal"
        :style="{ width: '25rem' }"
    >
        <span class="block mb-8">
            Please note that this will detach direct descendants from this
            family.
        </span>

        <div class="flex justify-between gap-2">
            <Button
                type="button"
                label="Cancel"
                severity="secondary"
                @click="cancel"
            ></Button>
            <Button
                type="button"
                label="Remove"
                severity="danger"
                @click="confirm"
            ></Button>
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
        saveFileStore.disableDeleteFamilyMode();
    }

    function confirm() {
        saveFileStore.disableDeleteFamilyMode(true);
    }
</script>
