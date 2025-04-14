

export async function getResetStream(baseURL=import.meta.env.VITE_API_URL) {
    return fetch(`${baseURL}/reset`)
  }
  
export async function getPreviousSegment(baseURL=import.meta.env.VITE_API_URL) {
    return fetch(`${baseURL}/previous`)
  }
