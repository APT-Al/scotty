(function () {
    'use strict'
  
    document.querySelector('[data-bs-toggle="offcanvas"]').addEventListener('click', function () {
      document.querySelector('.dashboard-collapse').classList.toggle('open')
    })
})()