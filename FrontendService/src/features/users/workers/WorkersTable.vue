<template>
  
  <div>
    <ModalOpener modalBoxId="addWorkerModal">
      <Button @click="handleAddWorkerClick" class="pull-right">Add Worker</Button>
    </ModalOpener>

    <Table>
      <TableHead :columnNames="['Name', 'Email', 'Sold Tickets', '']"/>
      <TableRow
        v-for="w in workers"
        :key="w.Id"
        :values="getWorkerValues(w)"
      >
        <ModalOpener modalBoxId="deleteWorkerModal">
          <DropdownMenu>
            <DropdownItem @click="handleViewWorkerClick(w)">View</DropdownItem>
            <DropdownItem @click="handleDeleteWorkerClick(w)">Delete</DropdownItem>
          </DropdownMenu>
        </ModalOpener>
      </TableRow>

    </Table>
  
    <Modal
      title="Add Worker"
      modalBoxId="addWorkerModal"
    >
        <div slot="body">
          <RegistrationForm
            :role=1
          />
      </div>
    </Modal>

    <Modal
      modalBoxId="deleteWorkerModal"
      title="Delete Worker"
    >
      <div slot="body">
        <p>Are you sure that you want to delete {{this.formWorkerName(this.selectedWorker)}} </p>
      </div>
      <div slot="buttons">
        <OptionModalButtons @yes="handleDeleteWorker" />
      </div>
    </Modal>
  </div>

  
</template>

<script>
import Table from '../../../components/Table/Table.vue';
import TableHead from '../../../components/Table/TableHead.vue';
import TableRow from '../../../components/Table/TableRow.vue';
import DropdownMenu from '../../../components/DropdownMenu/DropdownMenu.vue';
import DropdownItem from '../../../components/DropdownMenu/DropdownItem.vue';
import ModalOpener from '../../../components/Modal/ModalOpener.vue';
import Modal from '../../../components/Modal/Modal.vue';
import OptionModalButtons from '../../../components/Modal/OptionModalButtons.vue';
import { mapActions, mapGetters } from 'vuex';
import toastr from 'toastr'
import Button from '../../../components/Form/Button.vue';
import RegistrationForm from '../shared/RegistrationForm.vue';

export default {
  components: { Table, TableHead, TableRow, DropdownMenu, DropdownItem, ModalOpener, Modal, OptionModalButtons, Button, RegistrationForm },

  props: ['workers'],

  data: () => ({
    selectedWorker: null
  }),

  computed: {
    ...mapGetters({
      result: 'users/getResult'
    })
  },

  watch: {
    result({ message, ok, label }) {
      
      if (label === 'delete') {
        if (ok) {
          toastr.success(message);
        } else {
          toastr.error(message);
        }
      }
    }
  },
  
  methods: {
    ...mapActions({
      deleteWorker: 'users/deleteUser'
    }),
    
    formWorkerName(worker) {
      return `${worker?.Name} ${worker?.LastName}`;
    },

    getWorkerValues(worker) {
      return [
        this.formWorkerName(worker),
        worker.Email,
        worker.NumberOfSoldTickets
      ]
    },

    handleAddWorkerClick() {
      this.selectedWorker = null;
    },

    handleViewWorkerClick(w) {
      this.$router.push(`/users/${w.Id}`)
    },

    handleDeleteWorkerClick(w) {
      this.selectedWorker = w
    },

    handleDeleteWorker() {
      this.deleteWorker(this.selectedWorker?.Id);
    }
  }
}
</script>

<style scoped>

</style>