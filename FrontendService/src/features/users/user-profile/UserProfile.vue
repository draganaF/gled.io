<template>

  <div class="col-8">
    <Card title="User Profile">

      <PersonalInformationsForm
        :existingUser="user"
        :isEdit="isEdit"
      />

      <ChangePasswordForm
        v-if="isEdit"
        :existingUser="user"
      />

      <AdditionalUserInfoForm
        v-if="user && user.Role === 0"
        :user="user"
      />

      <AdditionalWorkerInfoForm
        v-if="user && user.Role === 1"
        :user="user"
      />

    </Card>
  </div>

</template>

<script>
import { mapActions, mapGetters } from 'vuex';

import Card from '../../../components/Card/Card.vue';
import PersonalInformationsForm from './PersonalInformationsForm.vue';
import ChangePasswordForm from './ChangePasswordForm.vue';
import AdditionalUserInfoForm from './AdditionalUserInfoForm.vue';
import AdditionalWorkerInfoForm from './AdditionalWorkerInfoForm.vue';
import { getUserIdFromToken } from '../../../utils/token';

export default {
  components: {
    Card,
    PersonalInformationsForm,
    ChangePasswordForm,
    AdditionalUserInfoForm,
    AdditionalWorkerInfoForm
  },

  data: () => ({
    isEdit: false
  }),

  computed: {
    ...mapGetters({
      user: 'users/getUser'
    })
  },

  methods: {
    ...mapActions({
      fetchUser: 'users/fetchUser'
    })
  },

  mounted() {
    let routeId = this.$route.params.id;
    if (routeId === "profile") {

      this.isEdit = true;
      routeId = getUserIdFromToken();
      console.log(routeId);
    }

    this.fetchUser(routeId);

    console.log("User profile page has mounted")
  }
}
</script>

<style>
</style>
