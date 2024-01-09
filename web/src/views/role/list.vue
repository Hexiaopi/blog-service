<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input v-model="listQuery.name" placeholder="名称" style="width: 200px;" class="filter-item"
        @keyup.enter.native="handleFilter" />
      <el-select v-model="listQuery.state" placeholder="状态" style="width: 90px" class="filter-item"
        @change="handleFilter">
        <el-option v-for="item in stateOptions" :key="item" :label="item | statusDisplayFilter" :value="item" />
      </el-select>
      <el-select v-model="listQuery.sort" style="width: 140px" class="filter-item" @change="handleFilter">
        <el-option v-for="item in sortOptions" :key="item.key" :label="item.label" :value="item.key" />
      </el-select>
      <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">
        搜索
      </el-button>
      <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleCreate">
        添加
      </el-button>
    </div>
    <el-table v-loading="listLoading" :data="list" element-loading-text="Loading" board fit highlight-current-row
      style="width: 100%;" @sort-change="sortChange">
      <el-table-column align="center" label="ID" min-width="50px">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="角色名称" align="center" width="100px">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="角色描述" align="center" min-width="100px">
        <template slot-scope="scope">
          {{ scope.row.desc }}
        </template>
      </el-table-column>
      <el-table-column label="状态" align="center" width="100">
        <template slot-scope="scope">
          <el-tag :type="scope.row.state | statusTypeFilter">{{ scope.row.state | statusDisplayFilter }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="create_time" label="创建时间" min-width="100px">
        <template slot-scope="scope">
          {{ scope.row.create_time }}
        </template>
      </el-table-column>
      <el-table-column align="center" prop="update_time" label="更新时间" min-width="100px">
        <template slot-scope="scope">
          {{ scope.row.update_time }}
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="230" class-name="small-padding fixed-width">
        <template slot-scope="{row,$index}">
          <!-- <el-button type="primary" size="mini" @click="handleUpdate(row)">
            编辑
          </el-button> -->
          <el-button type="primary" size="mini" @click="handleMenu(row)">
            菜单权限
          </el-button>
          <el-button type="primary" size="mini" @click="handleRest(row)">
            接口权限
          </el-button>
          <el-button v-if="row.state != 2" size="mini" type="danger" @click="handleDelete(row, $index)">
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total > 0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit"
      @pagination="getList" />


    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
      <el-form ref="dataForm" :rules="rules" :model="temp" label-position="left" label-width="70px"
        style="width: 400px; margin-left:50px;">
        <el-form-item label="角色名称" prop="title">
          <el-input v-model="temp.name" />
        </el-form-item>
        <el-form-item label="角色描述">
          <el-input v-model="temp.desc" :autosize="{ minRows: 2, maxRows: 4 }" type="textarea" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="temp.state" class="filter-item" placeholder="请选择">
            <el-option v-for="item in stateOptions" :key="item" :label="item | statusDisplayFilter" :value="item" />
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">
          取消
        </el-button>
        <el-button type="primary" @click="dialogStatus === 'create' ? createData() : updateData()">
          确认
        </el-button>
      </div>
    </el-dialog>

    <el-dialog title="菜单权限" :visible.sync="dialogMenuVisible">
      <el-form ref="menuForm" :model="temp" label-position="left" label-width="70px"
        style="width: 400px; margin-left:50px;">
        <el-tree :data="menus" ref="menuTree" show-checkbox node-key="id" :default-checked-keys="temp.menu_ids"
          default-expand-all :expand-on-click-node="false" :props="defaultMenuProps">
        </el-tree>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogMenuVisible = false">
          取消
        </el-button>
        <el-button type="primary" @click="updateRoleMenu()">
          确认
        </el-button>
      </div>
    </el-dialog>

    <el-dialog title="接口权限" :visible.sync="dialogRestVisible">
      <el-form ref="restForm" :model="temp" label-position="left" label-width="70px"
        style="width: 400px; margin-left:50px;">
        <el-tree :data="rests" ref="restTree" show-checkbox node-key="id" :default-checked-keys="temp.rest_ids"
          default-expand-all :expand-on-click-node="false" :props="defaultRestProps">
        </el-tree>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogRestVisible = false">
          取消
        </el-button>
        <el-button type="primary" @click="updateRoleRest()">
          确认
        </el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { listRole, createRole, updateRole, updateRoleMenu, updateRoleRest, deleteRole } from '@/api/role'
import { listMenu } from '@/api/menu'
import { listRest } from '@/api/rest'
import waves from '@/directive/waves' // waves directive
import Pagination from '@/components/Pagination' // secondary package based on el-pagination

export default {
  name: "Role",
  components: { Pagination },
  directives: { waves },
  filters: {
    statusTypeFilter(status) {
      const statusMap = {
        1: 'success',
        0: 'gray',
        2: 'danger'
      }
      return statusMap[status]
    },
    statusDisplayFilter(status) {
      const statusMap = {
        1: '有效',
        0: '无效'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 10,
        name: undefined,
        state: 1,
        sort: '+id'
      },
      stateOptions: [0, 1],
      sortOptions: [{ label: 'ID升序', key: '+id' }, { label: 'ID降序', key: '-id' }],
      temp: {
        id: undefined,
        name: '',
        desc: '',
        state: 1,
        menu_ids: undefined,
      },
      dialogFormVisible: false,
      menus: [],
      dialogMenuVisible: false,
      defaultMenuProps: {
        children: 'children',
        label: 'title'
      },
      rests: [],
      dialogRestVisible: false,
      defaultRestProps: {
        children: 'children',
        label: 'name'
      },
      dialogStatus: '',
      textMap: {
        update: '编辑',
        create: '创建'
      },
      rules: {
        name: [{ required: true, message: 'name is required', trigger: 'change' }],
        state: [{ required: true, message: 'state is required', trigger: 'blur' }]
      },


    }
  },
  created() {
    this.getList()
    this.getMenuList()
    this.getRestList()
  },
  methods: {
    getList() {
      this.listLoading = true
      listRole(this.listQuery).then(response => {
        this.list = response.data
        this.total = response.total
        this.listLoading = false
      })
    },
    getMenuList() {
      listMenu().then(response => {
        this.menus = response.data
      })
    },
    getRestList() {
      listRest().then(response => {
        this.rests = response.data
      })
    },
    handleFilter() {
      this.listQuery.page = 1
      this.getList()
    },
    sortChange(data) {
      const { prop, order } = data
      if (prop === 'id') {
        this.sortByID(order)
      }
    },
    sortByID(order) {
      if (order === 'ascending') {
        this.listQuery.sort = '+id'
      } else {
        this.listQuery.sort = '-id'
      }
      this.handleFilter()
    },
    resetTemp() {
      this.temp = {
        id: undefined,
        name: '',
        desc: '',
        state: 1,
        menu_ids: undefined,
      }
    },
    handleCreate() {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    createData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          createRole(this.temp).then(() => {
            this.dialogFormVisible = false
            this.$notify({
              title: 'Success',
              message: '创建成功',
              type: 'success',
              duration: 2000
            })
            this.getList();
          })
        }
      })
    },
    handleUpdate(row) {
      this.resetTemp()
      this.temp = Object.assign({}, row)
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    updateData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          const tempData = Object.assign({}, this.temp)
          updateRole(tempData).then(() => {
            const index = this.list.findIndex(v => v.id === this.temp.id)
            this.list.splice(index, 1, this.temp)
            this.dialogFormVisible = false
            this.$notify({
              title: 'Success',
              message: '更新成功',
              type: 'success',
              duration: 2000
            })
          })
        }
      })
    },
    handleMenu(row) {
      this.resetTemp()
      this.temp = Object.assign({}, row)
      this.dialogMenuVisible = true
      this.$nextTick(() => {
        this.$refs['menuForm'].clearValidate()
        //this.$refs.menuTree.setCheckedKeys(row.menu_ids)
      })
    },
    updateRoleMenu() {
      const menuIds = this.$refs.menuTree.getCheckedKeys()
      updateRoleMenu(this.temp.id, { "menu_ids": menuIds }).then(() => {
        this.dialogMenuVisible = false
        this.$notify({
          title: 'Success',
          message: '更新成功',
          type: 'success',
          duration: 2000
        })
        this.getList()
      })
    },
    handleRest(row) {
      this.resetTemp()
      this.temp = Object.assign({}, row)
      this.dialogRestVisible = true
      this.$nextTick(() => {
        this.$refs['restForm'].clearValidate()
        //this.$refs.restTree.setCheckedKeys(row.rest_ids)
      })
    },
    updateRoleRest() {
      const restIds = this.$refs.restTree.getCheckedKeys()
      updateRoleRest(this.temp.id, { "rest_ids": restIds }).then(() => {
        this.dialogRestVisible = false
        this.$notify({ title: 'Success', message: '更新成功', type: 'success', duration: 2000 })
      })
      this.getList()
    },
    handleDelete(row, index) {
      deleteRole(row.id).then(() => {
        this.$notify({ title: 'Success', message: '删除成功', type: 'success', duration: 2000 })
        this.list.splice(index, 1)
      })
    },
  }
}
</script>
