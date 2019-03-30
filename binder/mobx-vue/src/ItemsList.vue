<template>
    <ul>
        <li v-for="item in state.items"><a @click="state.removeItem(item)">[x]</a> {{ item }}</li>
    </ul>
</template>
<script lang="ts">
    import { observer } from "mobx-vue";
    export default observer({
        props: ['state']
    })
</script>