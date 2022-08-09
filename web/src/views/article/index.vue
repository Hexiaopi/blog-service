<template>
  <div class="app-container">

    <el-table v-loading="listLoading" :data="list" element-loading-text="Loading" stripe fit highlight-current-row>
      <el-table-column label="ID" width="95" align="center">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column label="标题" width="200" align="center">
        <template slot-scope="scope">
          {{ scope.row.title }}
        </template>
      </el-table-column>
      <el-table-column label="描述" width="200" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.desc }}</span>
        </template>
      </el-table-column>
      <el-table-column label="内容" width="200" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.content }}</span>
        </template>
      </el-table-column>
      <el-table-column label="图片" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.cover_image_url }}
        </template>
      </el-table-column>
      <el-table-column label="状态" width="110" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.state | statusTypeFilter">{{ scope.row.state | statusDisplayFilter }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="tag" label="文章标签" width="200" align="center">
        <template slot-scope="scope">
          <el-tag type="success">{{ scope.row.tags }}</el-tag>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { listArticle } from '@/api/article'

export default {
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
        1: '有效',
        0: '删除',
        2: '无效'
      }
      return statusMap[status]
    }
  },
  data () {
    return {
      list: null,
      listLoading: true
    }
  },
  created () {
    this.fetchData()
  },
  methods: {
    fetchData () {
      this.listLoading = true
      listArticle({ state: 1 }).then(response => {
        this.list = response.data
        this.listLoading = false
      })
    }
  }
}
</script>
