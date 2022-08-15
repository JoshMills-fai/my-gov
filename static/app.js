console.log('js loaded')

const pctCollection = document.querySelectorAll('.percentage');

pctCollection.forEach((item)=>{
    const inner = item.querySelector('.inner')
    
    const pct = inner.getAttribute('data-pct')
    console.log(pct)
    inner.style.width = pct + '%';
});
