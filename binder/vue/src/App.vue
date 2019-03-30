<template>
    <div>
        <h3>Todo: {{ items.length }}</h3>
        <ul>
            <li v-for="item in items">{{ item }}</li>
        </ul>
        <input type="text" v-model="newItem"/>
        <button @click="addNewItem">add</button>
    </div>
</template>
<script lang="ts">
    export default {
        data() {
            return {
                items: [],
                newItem: ''
            }
        },
        methods: {
            addNewItem() {
                this.items.push(this.newItem)
                this.newItem = ''
            }
        }
    }
</script>