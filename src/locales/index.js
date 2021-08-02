import {
    waitLocale,
    addMessages,
    register
} from 'svelte-i18n'

import zhCN from './zh-CN'
import enUS from './en-US'

export function initLoad() {
    addMessages('zh-CN', zhCN)
    addMessages('en-US', enUS)
}

export function SyncLoadLocale(name = 'zh-CN', locale = {}) {
    // import enUS from './en-US.json';
    // addMessages('en-US', enUS);
    addMessages(name, locale)
}

export function AsyncLoadLocale(name = 'zh-CN', locale = () => {}) {
    // register('en-US', () => import('./en-US.json'));
    register(name, locale)
}
export async function preload() {
    // awaits for the loading of the 'en-US' and 'en' dictionaries
    return waitLocale();
}