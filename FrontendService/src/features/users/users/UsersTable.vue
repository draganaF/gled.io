<template>
  
  <div>

    <Table>
      <TableHead :columnNames="['Name', 'Email', 'Bought Tickets', 'Reserved Tickets', 'Blocked', '']"/>
      <TableRow
        v-for="u in users"
        :key="u.Id"
        :values="getUserValues(u)"
      >
        <ModalOpener
          modalBoxId="blockUserModal"
          v-if="!u.Blocked"
        >
          <DropdownMenu>
            <DropdownItem @click="handleViewUserClick(u)">View</DropdownItem>
            <DropdownItem @click="handleBlockUserClick(u)">Block</DropdownItem>
          </DropdownMenu>
        </ModalOpener>
      </TableRow>

    </Table>
  
    <Modal
      modalBoxId="blockUserModal"
      title="Block User"
    >
      <div slot="body">
        <p>Are you sure that you want to block {{this.formUserName(this.selectedUser)}} </p>
      </div>
      <div slot="buttons">
        <OptionModalButtons @yes="handleBlockUser" />
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

export default {
  components: { Table, TableHead, TableRow, DropdownMenu, DropdownItem, ModalOpener, Modal, OptionModalButtons },

  props: ['users'],

  data: () => ({
    selectedUser: null
  }),

  computed: {
    ...mapGetters({
      result: 'users/getResult'
    })
  },

  watch: {
    result({ message, ok, label }) {
      
      if (label === 'block') {
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
      blockUser: 'users/blockUser'
    }),
    
    formUserName(user) {
      return `${user?.Name} ${user?.LastName}`;
    },

    getUserValues(user) {
      return [
        this.formUserName(user),
        user.Email,
        user.NumberOfBoughtTickets,
        user.NumberOfReservedTickets,
        user.Blocked ? 'Yes' : 'No'
      ]
    },

    handleViewUserClick(u) {
      this.$router.push(`/users/${u.Id}`)
    },

    handleBlockUserClick(u) {
      this.selectedUser = u
    },

    handleBlockUser() {
      this.blockUser(this.selectedUser?.Id);
    }
  }
}
</script>

<style scoped>

</style>