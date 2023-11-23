package product

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/christiandwi/edot/product-service/domain/products"
	"github.com/christiandwi/edot/product-service/entity"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
)

type service struct {
	productRepo products.ProductsRepository
}

func NewProductService(productRepo products.ProductsRepository) Service {
	return &service{
		productRepo: productRepo,
	}
}

func (serv *service) GetProducts() (err error) {
	// create chrome instance
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"),
		chromedp.WindowSize(1920, 1080),
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(
		ctx,
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// create a timeout
	// ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	// defer cancel()

	var nodes1 []*cdp.Node
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.tokopedia.com/p/handphone-tablet/handphone"),
		chromedp.KeyEvent(kb.End),
		chromedp.Sleep(15*time.Second),
		chromedp.Nodes(".css-bk6tzz", &nodes1, chromedp.ByQueryAll),
	); err != nil {
		panic(err)
	}

	timeScrapped := time.Now().Format("2006-01-02 15:04:05")

	var productName, imageLink, price, merchantName, productDetail string
	var count = 0
	var product entity.Products
	var products []entity.Products

	for i, node := range nodes1 {
		var rating = 0
		var ratingNodes []*cdp.Node
		if err = chromedp.Run(ctx,
			chromedp.AttributeValue("img", "src", &imageLink, nil, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(".css-20kt3o", &productName, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(".css-pp6b3e", &price, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(".css-vbihp9 > .css-ywdpwd:nth-child(2)", &merchantName, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Nodes(".css-177n1u3", &ratingNodes, chromedp.ByQueryAll, chromedp.FromNode(node)),
			// chromedp.Text(".css-17zm3l > div", &productDescription, chromedp.ByQuery),
			// chromedp.ActionFunc(func(ctx context.Context) error {
			// 	log.Printf("got description %v", productDescription)
			// 	return nil
			// }),
		); err != nil {
			panic(err)
		}

		// _, err := chromedp.RunResponse(ctx,
		// 	chromedp.Click(".css-16vw0vn", chromedp.ByQuery, chromedp.FromNode(node)),
		// )
		// if err != nil {
		// 	panic(err)
		// }

		// if err = chromedp.Run(ctx,
		// 	chromedp.Location(&productDetail),
		// 	chromedp.ActionFunc(func(ctx context.Context) error {
		// 		log.Printf("navigated to %v", productDetail)
		// 		return nil
		// 	}),
		// ); err != nil {
		// 	panic(err)
		// }

		for _, ratingNode := range ratingNodes {
			if ratingNode.AttributeValue("src") == "https://assets.tokopedia.net/assets-tokopedia-lite/v2/zeus/kratos/4fede911.svg" {
				rating++
			}
		}

		log.Printf("product number %v", i+1)
		log.Printf("got product name %v", productName)
		log.Printf("got image link %v", imageLink)
		log.Printf("got price %v", price)
		log.Printf("got merchant name %v", merchantName)
		log.Printf("got rating %v", rating)
		log.Printf("got product link %v", productDetail)
		// log.Printf("got description %v", productDescription)
		log.Println("======================================")

		product.ProductName = productName
		product.ImageLink = imageLink
		product.Price = price
		product.MerchantName = merchantName
		product.Rating = float32(rating)
		product.TimeScrapped = timeScrapped
		if err = serv.productRepo.InsertProducts(product); err != nil {
			panic(err)
		}
		products = append(products, product)
		count++
	}

	// if err := chromedp.Run(ctx,
	// 	chromedp.Navigate("https://www.tokopedia.com/p/handphone-tablet/handphone?page=2"),
	// 	chromedp.KeyEvent(kb.End),
	// 	chromedp.Sleep(10*time.Second),
	// 	chromedp.Nodes(".css-bk6tzz", &nodes2, chromedp.ByQueryAll),
	// ); err != nil {
	// 	panic(err)
	// }

	// for i, node := range nodes2 {
	// 	if count == 100 {
	// 		break
	// 	}
	// 	var rating = 0
	// 	var ratingNodes []*cdp.Node
	// 	if err = chromedp.Run(ctx,
	// 		chromedp.AttributeValue("img", "src", &imageLink, nil, chromedp.ByQuery, chromedp.FromNode(node)),
	// 		chromedp.Text(".css-20kt3o", &productName, chromedp.ByQuery, chromedp.FromNode(node)),
	// 		chromedp.Text(".css-pp6b3e", &price, chromedp.ByQuery, chromedp.FromNode(node)),
	// 		chromedp.Text(".css-vbihp9 > .css-ywdpwd:nth-child(2)", &merchantName, chromedp.ByQuery, chromedp.FromNode(node)),
	// 		chromedp.Nodes(".css-177n1u3", &ratingNodes, chromedp.ByQueryAll, chromedp.FromNode(node)),
	// 	); err != nil {
	// 		panic(err)
	// 	}

	// 	for _, ratingNode := range ratingNodes {
	// 		if ratingNode.AttributeValue("src") == "https://assets.tokopedia.net/assets-tokopedia-lite/v2/zeus/kratos/4fede911.svg" {
	// 			rating++
	// 		}
	// 	}

	// 	log.Printf("product number %v", i+1)
	// 	log.Printf("got product name %v", productName)
	// 	log.Printf("got image link %v", imageLink)
	// 	log.Printf("got price %v", price)
	// 	log.Printf("got merchant name %v", merchantName)
	// 	log.Printf("got rating %v", rating)
	// 	log.Printf("got product link %v", productDetail)
	// 	// log.Printf("got description %v", productDescription)
	// 	log.Println("======================================")

	// 	product.ProductName = productName
	// 	product.ImageLink = imageLink
	// 	product.Price = price
	// 	product.MerchantName = merchantName
	// 	product.Rating = float32(rating)
	// 	product.TimeScrapped = timeScrapped
	// 	if err = serv.productRepo.InsertProducts(product); err != nil {
	// 		panic(err)
	// 	}
	// 	products = append(products, product)
	// 	count++
	// }

	file, err := os.Create("public/products.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	headers := []string{
		"name",
		"image_link",
		"price",
		"rating",
		"merchant_name",
		"description",
	}
	// writing the column headers
	writer.Write(headers)

	for _, product := range products {
		record := []string{
			product.ProductName,
			product.ImageLink,
			product.Price,
			fmt.Sprintf("%.2f", product.Rating),
			product.MerchantName,
			product.Description,
		}

		// writing a new CSV record
		writer.Write(record)
	}
	defer writer.Flush()

	return
}
