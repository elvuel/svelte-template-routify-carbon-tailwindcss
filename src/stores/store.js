import {
    writable,
    get
} from 'svelte/store'


class GenericStore {
    constructor(key = 'default') {
        this.store_key = key
        this.objects = writable({})
        this.subscribe = this.objects.subscribe
        this.set = this.objects.set
        this.update = this.objects.update
    }
}

let stores = {}
let defaultStore = new GenericStore()
stores['default'] = defaultStore

export function newStore(store) {
    if (stores[store]) return
    stores[store] = new GenericStore(store)
}

export function getStore(store) {
    return stores[store] || defaultStore
}

export const store = stores['default']