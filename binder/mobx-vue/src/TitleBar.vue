<template>
    <h3>Total: {{ state.itemsCount }}</h3>
</template>
<script lang="ts">
    import { observer } from "mobx-vue";
    export default observer({
        props: ['state']
    })
</script>