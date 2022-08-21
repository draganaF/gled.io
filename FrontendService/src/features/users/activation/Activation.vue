<template>

  <div class="col-8">
    <Card title="Activation">
    </Card>
  </div>

</template>

<script>
import { mapActions, mapGetters } from 'vuex';
import Card from '../../../components/Card/Card.vue';
import toastr from 'toastr'

export default {
  components: {
    Card,
  },

  computed: {
    ...mapGetters({
      result: 'users/getResult'
    })
  },

  watch: {
    result({ ok, message, label }) {
      if(label !== 'activate') return;

      if(ok) {
        toastr.success(message);
        this.$router.push("/auth")
      } else {
        toastr.error(message);
      }
    }
  },

  methods: {

    ...mapActions({
      activateUser: 'users/activateUser'
    }),

    handleActivation() {
      const link = this.$route.query.link;
      if (!link) return;

      console.log(link);
      this.activateUser(link);
    }
  },

  mounted() {
    this.handleActivation()
  }
}
</script>

<style>
</style>
