export interface Task {
  id: number
  content: string
  status: 'pending' | 'in_progress' | 'completed'
  due_date?: string
  created_at: string
  updated_at: string
}

export interface CreateTaskRequest {
  content: string
  status: 'pending' | 'in_progress' | 'completed'
  due_date?: string
}

export interface UpdateTaskRequest {
  content: string
  status: 'pending' | 'in_progress' | 'completed'
  due_date?: string
}

const API_BASE_URL = '/api'

export const api = {
  async getTasks(): Promise<{ tasks: Task[] }> {
    const response = await fetch(`${API_BASE_URL}/tasks`)
    if (!response.ok) {
      throw new Error('Failed to fetch tasks')
    }
    return response.json()
  },

  async getTask(id: number): Promise<Task> {
    const response = await fetch(`${API_BASE_URL}/tasks/${id}`)
    if (!response.ok) {
      throw new Error('Failed to fetch task')
    }
    return response.json()
  },

  async createTask(task: CreateTaskRequest): Promise<Task> {
    const response = await fetch(`${API_BASE_URL}/tasks`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(task),
    })
    if (!response.ok) {
      throw new Error('Failed to create task')
    }
    return response.json()
  },

  async updateTask(id: number, task: UpdateTaskRequest): Promise<Task> {
    const response = await fetch(`${API_BASE_URL}/tasks/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(task),
    })
    if (!response.ok) {
      throw new Error('Failed to update task')
    }
    return response.json()
  },

  async deleteTask(id: number): Promise<void> {
    const response = await fetch(`${API_BASE_URL}/tasks/${id}`, {
      method: 'DELETE',
    })
    if (!response.ok) {
      throw new Error('Failed to delete task')
    }
  },
}