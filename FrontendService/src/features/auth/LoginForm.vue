<template>
    <Form>
        <form-row>
            <div class="col-12">
                <text-input 
                    label="Username"
                    v-model="Email"
                    type="text"
                />
            </div>
        </form-row>

        <form-row>
            <div class="col-12">
                <text-input 
                    label="Password"
                    v-model="Password"
                    type="password"
                />
            </div>
        </form-row>

        <Button @click="handleLoginClick">Sign In</Button>

        <div class="row mt-4">
            <p class="text-left col-6">
                <router-link to="/users/registration">Registration</router-link>
            </p>
        </div>
    </Form>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'

import Button from '../../components/Form/Button.vue'
import Form from '../../components/Form/Form.vue'
import FormRow from '../../components/Form/FormRow.vue'
import TextInput from '../../components/Form/TextInput.vue'
import toastr from 'toastr'

export default {
   components: {
       Form,
       FormRow,
       TextInput,
       Button
    },

    data: function() {
        return {
            Email: '',
            Password: ''
        }
    },

    computed: {
        ...mapGetters({result: 'authentication/getResult'})
    },

    watch: {
        result({ok, message, label}) {
            if(label !== 'authenticate') 
                return;

            if(ok) {
                this.$router.push('/projections');
            } else {
                toastr.error(message)
            }
            
        }
    },

    methods: {
        ...mapActions({authenticate: 'authentication/authenticate'}),

        handleLoginClick() {
            const authenticateObject = {
                Email: this.Email,
                Password: this.Password
            };

            this.authenticate(authenticateObject);
        }
    }
}

</script>
