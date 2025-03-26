import { createPinia } from 'pinia';
import useAppStore from './modules/app/index';
import useUserStore from './modules/user/index';
import useTabBarStore from './modules/tab-bar/index';
import useDictStore from './modules/dict/index';

const pinia = createPinia();

export { useAppStore, useUserStore, useTabBarStore, useDictStore };
export default pinia;
