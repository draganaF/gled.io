<template>
    <div class="col-md-6 col-lg-6 col-sm-10">
        <card :title="'Sign in'">
            <LoginForm />
        </card>
        <div>
            <p class="text-center">
                <router-link to="/projections">Back to home page</router-link>
            </p>
        </div>
    </div>
</template>

<script>
import Card from '../../components/Card/Card.vue'
import LoginForm from './LoginForm.vue'

export default {
    components: {
        Card,
        LoginForm
    }
}
</script>
