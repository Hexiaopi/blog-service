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
    </div>
    <el-table v-loading="listLoading" :data="list" element-loading-text="Loading" stripe fit highlight-current-row>
      <el-table-column label="ID" width="95" align="center">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column label="标题" width="200" align="center">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="描述" width="200" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.desc }}</span>
        </template>
      </el-table-column>
      <!-- <el-table-column label="内容" width="200" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.content }}</span>
        </template>
      </el-table-column>
      <el-table-column label="图片" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.cover_image_url }}
        </template>
      </el-table-column> -->
      <el-table-column label="状态" width="110" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.state | statusTypeFilter">{{ scope.row.state | statusDisplayFilter }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="tag" label="文章标签" width="200" align="center">
        <template slot-scope="scope">
          <el-tag v-for="tag in scope.row.tags" :key="tag.id">{{ tag.name }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" width="200" align="center">
        <template slot-scope="scope">
          {{ scope.row.create_time }}
        </template>
      </el-table-column>
      <el-table-column label="更新时间" width="200" align="center">
        <template slot-scope="scope">
          {{ scope.row.update_time }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作" width="120">
        <template slot-scope="scope">
          <router-link :to="'/article/edit/' + scope.row.id">
            <el-button type="primary" size="small" icon="el-icon-edit">
              编辑
            </el-button>
          </router-link>
        </template>
      </el-table-column>
    </el-table>
    <pagination v-show="total > 0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit"
      @pagination="getList" />
  </div>
</template>

<script>
import { listArticle } from '@/api/article'
import waves from '@/directive/waves' // waves directive
import Pagination from '@/components/Pagination' // secondary package based on el-pagination

export default {
  name: "Article",
  components: { Pagination },
  directives: { waves },
  filters: {
    statusTypeFilter (status) {
      const statusMap = {
        1: 'success',
        2: 'danger',
        3: 'gray'
      }
      return statusMap[status]
    },
    statusDisplayFilter (status) {
      const statusMap = {
        1: '发布',
        2: '草稿',
        3: '删除'
      }
      return statusMap[status]
    }
  },
  data () {
    return {
      list: null,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 10,
        name: undefined,
        state: 1,
        sort: '+id'
      },
      stateOptions: [1, 2, 3],
      sortOptions: [{ label: 'ID升序', key: '+id' }, { label: 'ID降序', key: '-id' }],
    }
  },
  created () {
    this.getList()
  },
  methods: {
    getList () {
      this.listLoading = true
      listArticle(this.listQuery).then(response => {
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
  }
}
</script>
