<template>
  <div class="app-container">
    <div class="filter-container">
      <el-form :inline="true" :model="listQuery">
        <el-form-item label="用户名">
          <el-input v-model="listQuery.name" clearable placeholder="用户名" style="width: 200px;" class="filter-item"
            @keyup.enter.native="handleFilter" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="listQuery.state" placeholder="状态" style="width: 90px" class="filter-item"
            @change="handleFilter">
            <el-option v-for="item in stateOptions" :key="item" :label="item | statusDisplayFilter" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item label="排序">
          <el-select v-model="listQuery.sort" style="width: 140px" class="filter-item" @change="handleFilter">
            <el-option v-for="item in sortOptions" :key="item.key" :label="item.label" :value="item.key" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">
            搜索
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit"
            @click="handleCreate">
            添加
          </el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table v-loading="listLoading" :data="list" element-loading-text="Loading" border stripe board fit
      highlight-current-row style="width: 100%;" @sort-change="sortChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column align="center" prop="id" label="ID" min-width="50px">
      </el-table-column>
      <el-table-column prop="name" label="名称" align="center" width="100px">
      </el-table-column>
      <el-table-column label="头像" align="center" width="100px">
        <template slot-scope="scope">
          <el-image :src="scope.row.avatar" lazy></el-image>
        </template>
      </el-table-column>
      <el-table-column label="状态" align="center" width="100px">
        <template slot-scope="scope">
          <el-tag :type="scope.row.state | statusTypeFilter">{{ scope.row.state | statusDisplayFilter }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="角色" align="center" width="200px">
        <template slot-scope="scope">
          <el-tag v-for="role in scope.row.roles" :key="role.id" type="success">{{ role.name }}</el-tag>
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
      <el-table-column fixed="right" label="操作" align="center" width="150" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-tooltip content="编辑" effect="dark" placement="top">
            <el-button type="primary" icon="el-icon-edit" size="mini" @click="handleUpdate(scope.row)">
              编辑
            </el-button>
          </el-tooltip>
          <el-tooltip class="delete-popover" content="删除" effect="dark" placement="top">
            <el-popconfirm title="确定删除吗？" @onConfirm="handleDelete(scope.row)">
              <el-button slot="reference" icon="el-icon-delete" size="mini" type="danger">
                删除
              </el-button>
            </el-popconfirm>
          </el-tooltip>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total > 0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit"
      @pagination="getList" />


    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible" width="30%">
      <el-form ref="dataForm" :rules="rules" :model="temp" label-position="left" label-width="100px"
        style="width: 400px; margin-left:50px;">
        <el-form-item label="用户名称" prop="name">
          <el-input v-model.trim="temp.name" placeholder="用户名" />
        </el-form-item>
        <el-form-item :label="dialogStatus === 'create' ? '用户密码' : '重置密码'" prop="password">
          <el-input v-model.trim="temp.password" autocomplte="off" :type="passwordType" />
          <span class="show-pwd" @click="showPwd">
            <svg-icon :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'" />
          </span>
        </el-form-item>
        <el-form-item label="角色" prop="roleIds">
          <el-select v-model.trim="temp.roles" value-key="id" multiple placeholder="请选择角色" style="width:100%">
            <el-option v-for="item in roles" :key="item.id" :label="item.name" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item label="用户头像">
          <el-input v-model="temp.avatar" :autosize="{ minRows: 2, maxRows: 4 }" type="textarea" placeholder="请输入" />
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
  </div>
</template>

<script>
import { listUser, createUser, updateUser, deleteUser } from '@/api/user'
import { listRole } from '@/api/role'
import waves from '@/directive/waves' // waves directive
import Pagination from '@/components/Pagination' // secondary package based on el-pagination

export default {
  name: "User",
  components: { Pagination },
  directives: { waves },
  filters: {
    statusTypeFilter (status) {
      const statusMap = {
        1: 'success',
        0: 'gray',
        2: 'danger'
      }
      return statusMap[status]
    },
    statusDisplayFilter (status) {
      const statusMap = {
        1: '无效',
        0: '有效'
      }
      return statusMap[status]
    }
  },
  data () {
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
        password: '',
        avatar: '',
        state: 0,
        roles: [],
      },
      roles: [],
      passwordType: 'password',
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: '编辑',
        create: '创建'
      },
      rules: {
        name: [{ required: true, message: 'type is required', trigger: 'change' }],
        state: [{ required: true, message: 'title is required', trigger: 'blur' }]
      },
    }
  },
  created () {
    this.getList()
    this.getRoles()
  },
  methods: {
    getRoles () {
      listRole({}).then(response => {
        this.roles = response.data
      })
    },
    getList () {
      this.listLoading = true
      listUser(this.listQuery).then(response => {
        this.list = response.data
        this.total = response.total
        this.listLoading = false
      })
    },
    handleFilter () {
      this.listQuery.page = 1
      this.getList()
    },
    sortChange (data) {
      const { prop, order } = data
      if (prop === 'id') {
        this.sortByID(order)
      }
    },
    sortByID (order) {
      if (order === 'ascending') {
        this.listQuery.sort = '+id'
      } else {
        this.listQuery.sort = '-id'
      }
      this.handleFilter()
    },
    resetTemp () {
      this.temp = {
        id: undefined,
        name: '',
        desc: '',
        state: 0,
        roles: [],
      }
    },
    showPwd () {
      if (this.passwordType === 'password') {
        this.passwordType = ''
      } else {
        this.passwordType = 'password'
      }
    },
    handleCreate () {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    createData () {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          createUser(this.temp).then(() => {
            this.list.unshift(this.temp)
            this.dialogFormVisible = false
            this.$notify({
              title: 'Success',
              message: '创建成功',
              type: 'success',
              duration: 2000
            })
          })
        }
      })
    },
    handleUpdate (row) {
      this.temp = Object.assign({}, row) // copy obj
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    updateData () {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          const tempData = Object.assign({}, this.temp)
          updateUser(tempData).then(() => {
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
    handleDelete (row) {
      deleteUser(row.id).then(() => {
        this.$notify({
          title: 'Success',
          message: '删除成功',
          type: 'success',
          duration: 2000
        })
        this.getList()
      })
    },
  }
}
</script>
<style scoped>
.show-pwd {
  position: absolute;
  right: 10px;
  top: 3px;
  font-size: 16px;
  color: #889aa4;
  cursor: pointer;
  user-select: none;
}
</style>
