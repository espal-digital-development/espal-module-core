<template>
    <div>
        LOGIN PAGE
        <form @submit.prevent="login">
            <input type="text" placeholder="email" autocomplete="email" v-model="credentials.email" />
            <input type="password" placeholder="password" autocomplete="current-password" v-model="credentials.password" />
            <input type="checkbox" v-model="credentials.rememberMe" />
            <button type="submit">Login!</button>
        </form>
    </div>
</template>

<script>
import { authService } from '../../services';
import { ref } from 'vue';

export default {
    setup() {
        const credentials = ref({
            email: '',
            password: '',
            rememberMe: false
        });
        const login = () => authService.login(credentials.value);
        return { credentials, login };
    }
};
</script>
