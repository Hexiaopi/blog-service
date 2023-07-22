<template>
  <div class="app-container">
    <div class="filter-container">
      <el-form :inline="true" :model="listQuery">
        <el-form-item label="操作者">
          <el-input v-model="listQuery.username" clearable placeholder="用户名" style="width: 200px;" class="filter-item"
            @keyup.enter.native="handleFilter" />
        </el-form-item>
        <el-form-item label="动作">
          <el-select v-model="listQuery.action" placeholder="动作" style="width: 90px" class="filter-item"
            @change="handleFilter">
            <el-option v-for="item in actionOptions" :key="item" :label="item | actionFilter" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item label="对象">
          <el-select v-model="listQuery.object" placeholder="对象" style="width: 90px" class="filter-item"
            @change="handleFilter">
            <el-option v-for="item in objectOptions" :key="item" :label="item | objectFilter" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item label="结果">
          <el-select v-model="listQuery.result" placeholder="结果" @change="handleFilter">
            <el-option value="Success" label="成功"></el-option>
            <el-option value="Fail" label="失败"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="排序方式">
          <el-select v-model="listQuery.sort" style="width: 140px" class="filter-item" @change="handleFilter">
            <el-option v-for="item in sortOptions" :key="item.key" :label="item.label" :value="item.key" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">
            搜索
          </el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table v-loading="listLoading" :data="list" element-loading-text="Loading" border stripe board fit
      highlight-current-row style="width: 100%;" @sort-change="sortChange">
      <el-table-column align="center" label="ID" min-width="50px">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="用户" align="center" width="100px">
        <template slot-scope="scope">
          {{ scope.row.user.name }}
        </template>
      </el-table-column>
      <el-table-column label="用户代理" align="center" min-width="100px">
        <template slot-scope="scope">
          {{ scope.row.user_agent }}
        </template>
      </el-table-column>
      <el-table-column label="访问IP" align="center" min-width="100px">
        <template slot-scope="scope">
          {{ scope.row.ip }}
        </template>
      </el-table-column>
      <el-table-column label="动作" align="center" width="100">
        <template slot-scope="scope">
          <el-tag :type="scope.row.action | actionDisplayFilter">{{ scope.row.action | actionFilter }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="对象" align="center" min-width="100px">
        <template slot-scope="scope">
          {{ scope.row.object | objectFilter }}
        </template>
      </el-table-column>
      <el-table-column label="结果" align="center" min-width="100px">
        <template slot-scope="scope">
          <div v-if="scope.row.result === 'Success'"><el-tag type="success">成功</el-tag></div>
          <div v-if="scope.row.result === 'Fail'">
            <el-tooltip :content="scope.row.error" placement="top">
              <el-tag type="danger">失败</el-tag>
            </el-tooltip>
          </div>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="create_time" label="操作时间" min-width="100px">
        <template slot-scope="scope">
          {{ scope.row.create_time }}
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="230" class-name="small-padding fixed-width">
        <template slot-scope="{row,$index}">
          <el-button size="mini" type="danger" @click="handleDelete(row, $index)">
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total > 0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit"
      @pagination="getList" />
  </div>
</template>

<script>
import { listOperation, deleteOperation } from '@/api/operation'
import waves from '@/directive/waves' // waves directive
import Pagination from '@/components/Pagination' // secondary package based on el-pagination

export default {
  name: "Operation",
  components: { Pagination },
  directives: { waves },
  filters: {
    objectFilter (object) {
      const objectMap = {
        'article': '文章',
        'tag': '标签',
        'resource': '资源',
        'user': '用户',
        'role': '角色',
      }
      return objectMap[object]
    },
    actionFilter (action) {
      const actionMap = {
        'POST': '创建',
        'PUT': '修改',
        'DELETE': '删除'
      }
      return actionMap[action]
    },
    actionDisplayFilter (action) {
      const actionMap = {
        'POST': 'success',
        'PUT': 'gray',
        'DELETE': 'danger'
      }
      return actionMap[action]
    },
  },
  data () {
    return {
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 10,
        username: '',
        object: '',
        action: '',
        result: '',
        sort: '+id'
      },
      objectOptions: ['article', 'tag', 'resource', 'user'],
      actionOptions: ['POST', 'PUT', 'DELETE'],
      sortOptions: [{ label: 'ID升序', key: '+id' }, { label: 'ID降序', key: '-id' }],
      temp: {
        id: undefined,
        username: '',
        object: '',
        action: '',
        result: '',
      },
    }
  },
  created () {
    this.getList()
  },
  methods: {
    getList () {
      this.listLoading = true
      listOperation(this.listQuery).then(response => {
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
        object: '',
        action: '',
        username: '',
        result: '',
      }
    },
    handleDelete (row, index) {
      deleteOperation(row.id).then(() => {
        this.$notify({
          title: 'Success',
          message: '删除成功',
          type: 'success',
          duration: 2000
        })
        this.list.splice(index, 1)
      })
    },
  }
}
</script>
