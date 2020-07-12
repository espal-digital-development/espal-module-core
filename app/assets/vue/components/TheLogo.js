import { h } from 'vue';

export default h('img', { src: require('../../images/logo.png').default.replace('js/spa/images', 'i') });
