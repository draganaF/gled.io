<template>

  <div>
    <div class="card">
      <div class="card-body ">
        <h6 class="card-category">
          <Stars :numOfStars="recension.score" />
        </h6>
        <h4 class="card-title">
          {{!!recension && recension.message}}
        </h4>
      </div>
      <div class="card-footer">
        <div class="author">
          <i class="material-icons md-light">account_circle</i>
          <span>{{formUserName(!!recension && recension.user_id)}}</span>
        </div>
        <div v-if="+userId !== recension.user_id" class="stats ml-auto">
          <ModalOpener
            :modalBoxId="`reportRecensionModal_${recension.id}`"
          >
            <i class="material-icons">feedback</i>Report
          </ModalOpener>
        </div>
      </div>
    </div>

    <Modal
      :modalBoxId="`reportRecensionModal_${recension.id}`"
      title="Report Recension"
    >
      <div slot="body">
        <p>Are you sure that you want to report this recension?</p>
      </div>
      <div slot="buttons">
        <OptionModalButtons @yes="handleReportRecension" />
      </div>
    </Modal>
  </div>

</template>

<script>
import { mapActions, mapGetters } from 'vuex';
import Stars from '../../../components/Rating/Stars.vue';
import { getUserIdFromToken } from '../../../utils/token';
import ModalOpener from '../../../components/Modal/ModalOpener.vue';
import OptionModalButtons from '../../../components/Modal/OptionModalButtons.vue';
import Modal from '../../../components/Modal/Modal.vue';
import toastr from 'toastr'

export default {

  components: { Stars, ModalOpener, OptionModalButtons, Modal },

  props: ['recension'],

  data: () => ({
    userId: getUserIdFromToken()
  }),

  computed: {
    ...mapGetters({
      users: 'users/getUsers',
      reportResult: 'reports/getResult'
    })
  },

  watch: {
    reportResult({ ok, message, label }) {
      if(label !== 'create') return;

      if(ok) {
        toastr.success(message);
      } else {
        toastr.error(message);
      }
    }
  },

  methods: {
    ...mapActions({
      createReport: 'reports/createReport',
    }),

    formUserName(userId) {
      const user = this.users.find(u => u.Id == userId);
      if (!user) return ``;

      return `${user.Name} ${user.LastName}`;
    },

    handleReportRecension() {
      this.createReport({
        id: 0,
        deleted: false,
        recension_id: this.recension.id,
        user_id: +getUserIdFromToken()
      });
    }
  }

}
</script>

<style>
</style>