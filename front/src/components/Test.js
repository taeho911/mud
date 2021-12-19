function Test(props) {
  const sessionTest = e => {
    e.preventDefault()
    let formData = new FormData(e.target.form)
    let jsonData = Object.fromEntries(formData.entries())
    fetch('/api/test/session', {
      method: 'post',
      headers: {'Content-Type': 'application/jsoncharset=UTF-8'},
      body: JSON.stringify(jsonData)
    }).then(res => res.json()).then(data => {
      console.log(data)
    }).catch(err => {
      console.log(err)
    })
  }

  return (
    <main>
      <h1>Test</h1>
      <form>
        <input type='text' name='cmd' placeholder='CMD'></input><br />
        <input type='text' name='username' placeholder='Username'></input><br />
        <input type='password' name='password' placeholder='Password'></input><br />
        <button onClick={e => sessionTest(e)}>Session Test</button>
      </form>
    </main>
  )
}

export default Test
