<template>
    <div>
        <TitleBar :state="domainModel"/>
        <ItemsList :state="domainModel"/>
        <NewItemPanel :state="newItemPanelViewModel"/>
    </div>
</template>
<script lang="ts">
    import { observer } from "mobx-vue";
    import DomainModel from "./DomainModel.ts";
    import NewItemPanelViewModel from "./NewItemPanelViewModel.ts"
    import TitleBar from "./TitleBar";
    import ItemsList from "./ItemsList";
    import NewItemPanel from "./NewItemPanel"

    export default observer({
        components: {TitleBar, ItemsList, NewItemPanel},
        data() {
            let domainModel = new DomainModel();
            let newItemPanelViewModel = new NewItemPanelViewModel(domainModel);
            return { domainModel, newItemPanelViewModel }
        }
    })
</script>