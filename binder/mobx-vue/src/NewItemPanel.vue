<template>
    <div>
        <input type="text" v-model="state.newItem" />
        <button @click="state.addItem()">Add</button>
    </div>
</template>
<script lang="ts">
    import { observer } from "mobx-vue";
    export default observer({
        props: ['state']
    })
</script>