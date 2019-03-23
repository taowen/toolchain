[[toc]]

# 要解决的问题

HTML DOM, Mysql Table, ElaticSearch 等等技术都有一个自己对于状态的表现形式。
为了操作这些不同的表现形式，我们需要使用不同的API，不同的语法。
但是本质上来说，这些状态都可以用统一的抽象形式来表达。
当我们在表达状态，以及状态的变化的时候。可以把形式上的差异屏蔽掉，关注其本质，可以简化思维复杂度。

# 解决方案

从原表现形式，监听其改变事件，然后作用于目标表现形式。

## 构成

| 构成 | 解释 |
| --- | --- |
| source | 原表现形式 |
| target | 目标表现形式 |
| binder | 实现binding的中间人 |

# 解决方案案例

## mobx-vue

### DomainModel.ts
<<< @/binder/mobx-vue/src/DomainModel.ts

这里定义的是这个应用的业务逻辑。也就是其实质性的东西。而视图是对这份状态的一种呈现方式。

### App.vue
<<< @/binder/mobx-vue/src/App.vue

### TitleBar.vue
<<< @/binder/mobx-vue/src/TitleBar.vue

从 TitleBar 这里，我们可以看出来。View是对DomainModel做了一个单向的绑定。如果DomainModel更新了，View里的Total数量就会自动更新。

### ItemsList.vue
<<< @/binder/mobx-vue/src/ItemsList.vue

这里的 ItemsList 和 TitleBar 一样，都是一个视图，是对同一份数据的两份不同的呈现。一个展示了总数，一个展示了明细。都会被单向同步。

### NewItemPanelViewModel.ts
<<< @/binder/mobx-vue/src/NewItemPanelViewModel.ts

这里定义的是一个视图的私有的状态。它承上启下。
一方面和NewItemPanel这个视图进行了连接。
另外一方面和DomainModel进行了连接。
然后通过DomainModel，它又与TitleBar和ItemsList这两个视图进行了间接的连接关系。

### NewItemPanel.vue
<<< @/binder/mobx-vue/src/NewItemPanel.vue

这里的 NewItemPanel 和 NewItemPanelViewModel 的关系是双向绑定。其实就是用 ViewModel 来做为 View 的替身。
这样我们所有需要取的数据，不用管View了，直接从ViewModel上取。
我们所有对View希望作的更新，也不用更新View了，直接更新到ViewModel上。
两者之间的同步，由 vue 框架 `v-model="state.newItem"` 来保证。


| 构成 | 解释 |
| --- | --- |
| source | HTML DOM |
| target | Javascript Object Model |
| binder | vue + mobx |

## vue

### App.vue
<<< @/binder/vue/src/App.vue

vue 没有 mobx + vue 那么灵活。状态和视图必须在一个component里。如果要多个视图共享一个后端状态，就需要其他库的协助。

| 构成 | 解释 |
| --- | --- |
| source | HTML DOM |
| target | Javascript Object Model |
| binder | vue |