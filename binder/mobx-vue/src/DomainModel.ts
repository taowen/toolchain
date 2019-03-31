import { action, computed, observable } from "mobx";
export default class DomainModel {
    @observable items: string[] = [];

    @computed get itemsCount() {
        return this.items.length;
    }

    @action.bound addItem(newItem: string) {
        this.items.push(newItem);
    }

    @action.bound removeItem(item: string) {
        this.items.splice(this.items.indexOf(item), 1);
    }
}