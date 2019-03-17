import { action, computed, observable } from "mobx";
import DomainModel from "./DomainModel"
export default class NewItemPanelViewModel {
    @observable domainModel: DomainModel;
    @observable newItem = '';

    constructor(domainModel: DomainModel) {
        this.domainModel = domainModel;
    }

    @action.bound addItem() {
        this.domainModel.addItem(this.newItem);
        this.newItem = '';
    }
}