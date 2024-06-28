let selectElement = document.querySelector('.shop__product__option__right select');

selectElement.onchange = () => {
    // Get product elements
    let productElements = document.querySelectorAll('.col-lg-4.col-md-6.col-sm-6');

    // Convert NodeList to Array
    let products = Array.from(productElements).map((productElement, index) => {
        let rect = productElement.getBoundingClientRect();
        return {
            name: productElement.querySelector('.product__item__text h6 a').textContent,
            price: parseFloat(productElement.querySelector('.product__item__text h5').textContent.replace('$', '')),
            element: productElement,
            position: { top: rect.top, left: rect.left },
            index: index
        };
    });

    // Sort products
    if (selectElement.value === 'lth') {
        // Sort from low to high
        products.sort((a, b) => a.price - b.price);
    } else if (selectElement.value === 'htl') {
        // Sort from high to low
        products.sort((a, b) => b.price - a.price);
    }

    // Animate sorted products
    products.forEach((product, index) => {
        let rect = product.element.getBoundingClientRect();
        let translateX = product.position.left - rect.left;
        let translateY = product.position.top - rect.top;

        product.element.style.transform = `translate(${translateX}px, ${translateY}px)`;
        product.element.style.position = 'relative';
        product.element.style.zIndex = '1';

        requestAnimationFrame(() => {
            product.element.style.transition = 'transform 0.5s';
            product.element.style.transform = '';
        });

        setTimeout(() => {
            product.element.style.transition = '';
            product.element.style.position = '';
            product.element.style.zIndex = '';
        }, 500);
    });

    // Clear current products
    let catalog = document.querySelector('#catalog');
    while (catalog.firstChild) {
        catalog.removeChild(catalog.firstChild);
    }

    // Append sorted products to the catalog
    products.forEach(product => {
        catalog.appendChild(product.element);
    });
};

console.log(selectElement);


const button = document.querySelector(".addtocart");
const done = document.querySelector(".done");
console.log(button);
let added = false;
button.addEventListener('click',()=>{
    if(added){
        done.style.transform = "translate(-110%) skew(-40deg)";
        added = false;
    }
    else{
        done.style.transform = "translate(0px)";
        added = true;
    }

});
