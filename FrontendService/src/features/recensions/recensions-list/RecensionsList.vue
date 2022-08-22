<template>
  
  <div>

    <ModalOpener
      modalBoxId="giveRecensionModal"
      v-if="shouldShowButton()"
    >
      <Button class="pull-right">Give Recension</Button>
    </ModalOpener>

    <RecensionItem v-for="r in recensions"
      :key="r.id"
      :recension="r"
    />

    <Modal
      modalBoxId="giveRecensionModal"
      title="Give Recension"
    >
      <div slot="body">
        <GiveRecensionForm :movieId="movieId" />
      </div>
    </Modal>
  </div>

</template>

<script>
import { mapActions, mapGetters } from 'vuex'
import RecensionItem from './RecensionItem.vue';
import ModalOpener from '../../../components/Modal/ModalOpener.vue';
import Button from '../../../components/Form/Button.vue';
import Modal from '../../../components/Modal/Modal.vue';
import GiveRecensionForm from './GiveRecensionForm.vue';
import { getUserIdFromToken, getRoleFromToken } from '../../../utils/token';
import recensions from '../../../store/modules/recensions';

export default {
  
  components: { RecensionItem, ModalOpener, Button, Modal, GiveRecensionForm },

  props: ['movieId'],

  data: () => ({
    userId: getUserIdFromToken(),
    role: getRoleFromToken()
  }),

  computed: {
    ...mapGetters({
      recensions: 'recensions/getRecensions',
      result: 'recensions/getResult'
    })
  },

  watch: {
    movieId() {
      this.fetchRecensions();
    },

    result({ ok, label }) {
      if (!ok) return;

      label === 'create' && this.fetchRecensions();
    }
  },

  methods: {
    ...mapActions({
      fetchRecensionsForMovie: 'recensions/fetchByMovie',
      fetchUsers: 'users/searchUsers'
    }),

    shouldShowButton() {
      return !this.recensions || !this.recensions?.find(r => r.user_id == +this.userId); 
    },

    fetchRecensions() {
      this.fetchRecensionsForMovie(this.movieId);
    }
  },

  mounted() {
    this.fetchRecensions();
    this.fetchUsers({});
    console.log(getRoleFromToken())
  }
}
</script>
