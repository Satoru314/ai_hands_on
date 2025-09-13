import { useState, useEffect, useCallback } from 'react'
import { Task, CreateTaskRequest, UpdateTaskRequest, api } from '@/lib/api'

interface UseTasksReturn {
  tasks: Task[]
  loading: boolean
  error: string | null
  createTask: (taskData: CreateTaskRequest) => Promise<void>
  updateTask: (id: number, updates: Partial<UpdateTaskRequest>) => Promise<void>
  deleteTask: (id: number) => Promise<void>
  refreshTasks: () => Promise<void>
  clearError: () => void
}

export const useTasks = (): UseTasksReturn => {
  const [tasks, setTasks] = useState<Task[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  const clearError = useCallback(() => {
    setError(null)
  }, [])

  const refreshTasks = useCallback(async () => {
    try {
      setLoading(true)
      setError(null)
      const response = await api.getTasks()
      setTasks(response.tasks)
    } catch (err) {
      console.error('Failed to fetch tasks:', err)
      setError('タスクの取得に失敗しました')
    } finally {
      setLoading(false)
    }
  }, [])

  const createTask = useCallback(async (taskData: CreateTaskRequest) => {
    try {
      setError(null)
      const newTask = await api.createTask(taskData)
      setTasks(prev => [...prev, newTask])
    } catch (err) {
      console.error('Failed to create task:', err)
      setError('タスクの作成に失敗しました')
      throw err
    }
  }, [])

  const updateTask = useCallback(async (id: number, updates: Partial<UpdateTaskRequest>) => {
    try {
      setError(null)

      // 現在のタスクを取得
      const currentTask = tasks.find(task => task.id === id)
      if (!currentTask) {
        setError('タスクが見つかりません')
        throw new Error('Task not found')
      }

      // 完全な更新データを構築
      const fullUpdate: UpdateTaskRequest = {
        content: updates.content ?? currentTask.content,
        status: updates.status ?? currentTask.status,
        ...(updates.due_date !== undefined
          ? { due_date: updates.due_date }
          : currentTask.due_date
            ? { due_date: currentTask.due_date }
            : {})
      }

      console.log('Sending update request:', { id, fullUpdate })

      const updatedTask = await api.updateTask(id, fullUpdate)
      setTasks(prev => prev.map(task => task.id === id ? updatedTask : task))
    } catch (err) {
      console.error('Failed to update task:', err)
      setError('タスクの更新に失敗しました')
      throw err
    }
  }, [tasks])

  const deleteTask = useCallback(async (id: number) => {
    try {
      setError(null)
      await api.deleteTask(id)
      setTasks(prev => prev.filter(task => task.id !== id))
    } catch (err) {
      console.error('Failed to delete task:', err)
      setError('タスクの削除に失敗しました')
      throw err
    }
  }, [])

  // 初回ロード
  useEffect(() => {
    refreshTasks()
  }, [refreshTasks])

  return {
    tasks,
    loading,
    error,
    createTask,
    updateTask,
    deleteTask,
    refreshTasks,
    clearError,
  }
}