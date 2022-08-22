<template>
  
  <div>
    <Table>
      <TableHead :columnNames="['Message', 'Written By', 'Reported By', '']" />
      <TableRow
        v-for="r in reports"
        :key="r.id"
        :values="getReportValues(r)"
      >
        <DropdownMenu>
          <ModalOpener
            modalBoxId="deleteRecensionModal"
          >
            <DropdownItem @click="handleDeleteRecensionClick(r)">Delete Recension</DropdownItem>
          </ModalOpener>

          <ModalOpener
            modalBoxId="blockUserModal"
            v-if="showBlockUser(r)"
          >
            <DropdownItem @click="handleBlockUserClick(r)">Block User</DropdownItem>
          </ModalOpener>

          <ModalOpener
            modalBoxId="ignoreReportModal"
          >
            <DropdownItem @click="handleIgnoreReportClick(r)">Ignore Report</DropdownItem>
          </ModalOpener>

        </DropdownMenu>


      </TableRow>
    </Table>

    <Modal
      modalBoxId="ignoreReportModal"
      title="Ignore Rerpot"
    >
      <div slot="body">
        <p>Are you sure that you want to ignore this report</p>
      </div>
      <div slot="buttons">
        <OptionModalButtons @yes="handleIgnoreReport" />
      </div>
    </Modal>

    <Modal
      modalBoxId="blockUserModal"
      title="Block User"
    >
      <div slot="body">
        <p>Are you sure that you want to block user who wrote this recension?</p>
      </div>
      <div slot="buttons">
        <OptionModalButtons @yes="handleBlockUser" />
      </div>
    </Modal>

    <Modal
      modalBoxId="deleteRecensionModal"
      title="Delete Recension"
    >
      <div slot="body">
        <p>Are you sure that you want to delete this recension?</p>
      </div>
      <div slot="buttons">
        <OptionModalButtons @yes="handleDeleteRecension" />
      </div>
    </Modal>

  </div>
</template>

<script>
import TableHead from '../../../components/Table/TableHead.vue';
import TableRow from '../../../components/Table/TableRow.vue';
import Table from '../../../components/Table/Table.vue';
import { mapActions, mapGetters } from 'vuex';
import DropdownMenu from '../../../components/DropdownMenu/DropdownMenu.vue';
import ModalOpener from '../../../components/Modal/ModalOpener.vue';
import DropdownItem from '../../../components/DropdownMenu/DropdownItem.vue';
import Modal from '../../../components/Modal/Modal.vue';
import OptionModalButtons from '../../../components/Modal/OptionModalButtons.vue';
import toastr from 'toastr'

export default {

  components: { TableHead, TableRow, Table, DropdownMenu, ModalOpener, DropdownItem, Modal, OptionModalButtons },

  props: ['reports'],

  data: () => ({
    selectedReport: undefined
  }),

  computed: {
    ...mapGetters({
      users: 'users/getUsers',
      recensions: 'recensions/getRecensions',
      reportResult: 'reports/getResult',
      userResult: 'users/getResult',
      recensionResult: 'recensions/getResult'
    })
  },

  watch: {
    reportResult({ ok, message, label }) {
      if(label !== 'delete') return;

      if(ok) {
        toastr.success(message);
      } else {
        toastr.error(message);
      }
    },

    userResult({ message, ok, label }) {
      if (label !== 'block') return;

      if (ok) {
        toastr.success(message);
        this.fetchUsers({});
      } else {
        toastr.error(message);
      }
    },

    recensionResult({ message, ok, label }) {
      if (label !== 'delete') return;

      if (ok) {
        toastr.success(message);
        this.silentDeleteReport(this.selectedReport.id);
      } else {
        toastr.error(message);
      }
    }
  },

  methods: {
    ...mapActions({
      fetchUsers: 'users/searchUsers',
      fetchRecensions: 'recensions/fetchRecensions',
      deteleReport: 'reports/deleteReport',
      blockUser: 'users/blockUser',
      deleteRecension: 'recensions/deleteRecension',
      silentDeleteReport: 'reports/silentDeleteReport'
    }),

    getReportValues(report) {
      const recension = this.recensions?.find(r => r.id === report.recension_id);

      return [
        recension?.message ?? '',
        this.formUserName(recension?.user_id),
        this.formUserName(report.user_id)
      ]
    },

    formUserName(userId) {
      const user = this.users?.find(u => u.Id == userId);
      return `${user?.Name} ${user?.LastName} ${user && user?.Blocked ? '(Blocked)' : ''}`
    },

    showBlockUser(report) {
      const recension = this.recensions?.find(r => r.id === report.recension_id);
      const user = this.users?.find(u => u.Id == recension?.user_id);

      return !user?.Blocked;
    },

    handleIgnoreReportClick(report) {
      this.selectedReport = report;
    },

    handleIgnoreReport() {
      this.deteleReport(this.selectedReport?.id);
    },

    handleBlockUserClick(report) {
      this.selectedReport = report;
    },

    handleBlockUser() {
      const recension = this.recensions?.find(r => r.id === this.selectedReport.recension_id);
      this.blockUser(recension?.user_id);
    },

    handleDeleteRecensionClick(report) {
      this.selectedReport = report;
    },

    handleDeleteRecension() {
      this.deleteRecension(this.selectedReport?.recension_id);
    }
  },

  mounted() {
    this.fetchRecensions();
    this.fetchUsers({});
  }
}
</script>

<style scoped>
</style>