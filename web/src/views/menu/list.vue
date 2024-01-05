<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input v-model="listQuery.name" placeholder="名称" style="width: 200px;" class="filter-item"
        @keyup.enter.native="handleFilter" />
      <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">
        搜索
      </el-button>
      <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-plus"
        @click="handleCreate(0)">
        添加
      </el-button>
    </div>
    <el-table v-loading="listLoading" :data="list" element-loading-text="Loading" board fit highlight-current-row
      style="width: 100%;" row-key="id" :tree-props="{ children: 'children', hasChildren: 'hasChildren' }">
      <el-table-column label="名称" align="center" width="100px">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="标题" align="center" width="100px">
        <template slot-scope="scope">
          {{ scope.row.meta.title }}
        </template>
      </el-table-column>
      <el-table-column label="图标" align="center" width="100px">
        <template slot-scope="scope">
          <i :class="scope.row.meta.icon"></i>
        </template>
      </el-table-column>

      <el-table-column label="路径" align="center" min-width="100px">
        <template slot-scope="scope">
          {{ scope.row.path }}
        </template>
      </el-table-column>
      <el-table-column label="重定向" align="center" min-width="100px">
        <template slot-scope="scope">
          {{ scope.row.redirect }}
        </template>
      </el-table-column>
      <el-table-column label="组件" align="center" min-width="100px">
        <template slot-scope="scope">
          {{ scope.row.component }}
        </template>
      </el-table-column>
      <el-table-column label="顺序" align="center" min-width="100px">
        <template slot-scope="scope">
          {{ scope.row.order }}
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
          <el-button v-if="row.parent_id == 0" type="primary" icon="el-icon-plus" size="mini" @click="handleCreate(row.id)">
            新增
          </el-button>
          <el-button type="primary" icon="el-icon-edit" size="mini" @click="handleUpdate(row)">
            编辑
          </el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="handleDelete(row, $index)">
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total > 0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit"
      @pagination="getList" />


    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
      <el-form ref="dataForm" :rules="rules" :model="temp" label-position="left" label-width="70px"
        style="width: 500px; margin-left:50px;">
        <el-form-item label="名称">
          <el-input v-model="temp.name"></el-input>
        </el-form-item>
        <el-form-item label="路径">
          <el-input v-model="temp.path"></el-input>
        </el-form-item>
        <el-form-item label="重定向">
          <el-input v-model="temp.redirect"></el-input>
        </el-form-item>
        <el-form-item label="标题">
          <el-input v-model="temp.title"></el-input>
        </el-form-item>
        <el-form-item label="图标">
          <el-select v-model="temp.icon" filterable placeholder="请选择">
            <el-option v-for="item in elementIcons" :key="item" :label="item" :value="item">
              <i :class="item">{{ item }}</i>
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="组件">
          <el-input v-model="temp.component"></el-input>
        </el-form-item>
        <el-form-item label="顺序">
          <el-input-number v-model="temp.order"></el-input-number>
        </el-form-item>
        <el-form-item label="展示">
          <el-radio-group v-model="temp.show">
            <el-radio :label="0">不展示</el-radio>
            <el-radio :label="1">展示</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="隐藏">
          <el-radio-group v-model="temp.hidden">
            <el-radio :label="0">不隐藏</el-radio>
            <el-radio :label="1">隐藏</el-radio>
          </el-radio-group>
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
import { listMenuTree, createMenu, updateMenu, deleteMenu } from '@/api/menu'
import waves from '@/directive/waves' // waves directive
import Pagination from '@/components/Pagination' // secondary package based on el-pagination
import elementIcons from './element-icons'

export default {
  name: "Menu",
  components: { Pagination, elementIcons },
  directives: { waves },
  filters: {

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
      },
      temp: {
        id: 0,
        name: '',
        path: '',
        redirect: '',
        title: '',
        icon: '',
        component: '',
        order: 0,
        show: true,
        hidden: false,
        parent_id: 0,
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: '编辑',
        create: '创建'
      },
      rules: {
        name: [{ required: true, message: 'name is required', trigger: 'change' }],
        path: [{ required: true, message: 'path is required', trigger: 'blur' }],
      },
      elementIcons
    }
  },
  created() {
    this.getList()
  },
  methods: {
    generateElementIconCode(symbol) {
      return `<i class="el-icon-${symbol}" />`
    },
    getList() {
      this.listLoading = true
      listMenuTree(this.listQuery).then(response => {
        this.list = response.data
        this.total = response.total
        this.listLoading = false
      })
    },
    handleFilter() {
      this.listQuery.page = 1
      this.getList()
    },
    resetTemp() {
      this.temp = {
        id: 0,
        name: '',
        path: '',
        redirect: '',
        title: '',
        icon: '',
        component: '',
        order: 0,
        show: true,
        hidden: false,
        parent_id: 0,
      }
    },
    handleCreate(id) {
      this.resetTemp()
      if (id > 0) {
        this.temp.parent_id = id
      }
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    createData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          createMenu(this.temp).then(() => {
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
      this.temp = Object.assign({}, row) // copy obj
      this.temp.title = row.meta.title
      this.temp.icon = row.meta.icon
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
          updateMenu(tempData).then(() => {
            this.getList()
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
    handleDelete(row, index) {
      deleteMenu(row.id).then(() => {
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
