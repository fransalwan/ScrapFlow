package seeders

import (
	"log"
	"scrap-invoice-backend/config"
	"scrap-invoice-backend/models"
)

// Tambahin fungsi ini di paling atas atau bawah file seeder.go
func SeedData() {
	log.Println("🌱 Starting seeding process...")

	seedCustomers()
	seedItemCategories()
	seedItems()

	log.Println("🎉 Seeding process completed!")
}

func seedCustomers() {
	var count int64
	config.DB.Model(&models.Customer{}).Count(&count)
	if count > 0 {
		log.Println("⏭️  Customers already seeded, skipping...")
		return
	}

	customers := []models.Customer{
		{
			Name:    "PT. Baja Jaya",
			Email:   "ptbaja@example.com",
			Phone:   "081234567890",
			Address: "Jl. Besi No. 1",
			Tier:    "Gold",
		},
		{
			Name:    "CV. Efrata",
			Email:   "efrata@example.com",
			Phone:   "085678123456",
			Address: "Jl. Baja Mulia No. 12",
			Tier:    "Silver",
		},
		{
			Name:    "UD. Logam Makmur",
			Email:   "logam@example.com",
			Phone:   "087812345678",
			Address: "Jl. Besi Lama No. 21",
			Tier:    "Gold",
		},
		{
			Name:    "PT. Mandiri Scrap",
			Email:   "mandiri@example.com",
			Phone:   "089912345678",
			Address: "Jl. Baja Tua No. 3",
			Tier:    "Bronze",
		},
	}

	for _, customer := range customers {
		config.DB.Create(&customer)
	}
	log.Println("✅ Customers seeded successfully (4 customers)")
}

func seedItemCategories() {
	var count int64
	config.DB.Model(&models.ItemCategory{}).Count(&count)
	if count > 0 {
		log.Println("⏭️  Item categories already seeded, skipping...")
		return
	}

	categories := []models.ItemCategory{
		{ItemCategoryName: "Besi"},
		{ItemCategoryName: "Tembaga"},
		{ItemCategoryName: "Aluminium"},
		{ItemCategoryName: "Plastik"},
		{ItemCategoryName: "Kuningan"},
		{ItemCategoryName: "Seng"},
		{ItemCategoryName: "Timah"},
		{ItemCategoryName: "Akunium"},
		{ItemCategoryName: "Kaleng"},
		{ItemCategoryName: "Kabel"},
	}

	for _, category := range categories {
		config.DB.Create(&category)
	}
	log.Println("✅ Item categories seeded successfully (10 categories)")
}

func seedItems() {
	var count int64
	config.DB.Model(&models.Item{}).Count(&count)
	if count > 0 {
		log.Println("⏭️  Items already seeded, skipping...")
		return
	}

	// Ambil semua categories
	var categories []models.ItemCategory
	if err := config.DB.Find(&categories).Error; err != nil {
		log.Fatalf("Failed to fetch categories: %v", err)
	}

	// Buat map untuk nyari category by name
	categoriesMap := make(map[string]models.ItemCategory)
	for _, cat := range categories {
		categoriesMap[cat.ItemCategoryName] = cat
	}

	items := []models.Item{
		// Besi
		{
			ItemName:       "Besi Beton",
			ItemCategoryID: categoriesMap["Besi"].ID,
			PricePerKg:     8800.00,
		},
		// Tembaga
		{
			ItemName:       "Tembaga Kupas",
			ItemCategoryID: categoriesMap["Tembaga"].ID,
			PricePerKg:     94000.00,
		},
		{
			ItemName:       "Tembaga Bakar",
			ItemCategoryID: categoriesMap["Tembaga"].ID,
			PricePerKg:     92000.00,
		},
		// Aluminium
		{
			ItemName:       "Aluminium Kaleng",
			ItemCategoryID: categoriesMap["Aluminium"].ID,
			PricePerKg:     12000.00,
		},
		{
			ItemName:       "Aluminium Piringan",
			ItemCategoryID: categoriesMap["Aluminium"].ID,
			PricePerKg:     11000.00,
		},
		// Plastik
		{
			ItemName:       "Plastik Botol",
			ItemCategoryID: categoriesMap["Plastik"].ID,
			PricePerKg:     3500.00,
		},
		{
			ItemName:       "Plastik Campur",
			ItemCategoryID: categoriesMap["Plastik"].ID,
			PricePerKg:     3000.00,
		},
		// Kuningan
		{
			ItemName:       "Kuningan Kuning",
			ItemCategoryID: categoriesMap["Kuningan"].ID,
			PricePerKg:     58000.00,
		},
		{
			ItemName:       "Kuningan Merah",
			ItemCategoryID: categoriesMap["Kuningan"].ID,
			PricePerKg:     56000.00,
		},
		// Seng
		{
			ItemName:       "Seng Gelombang",
			ItemCategoryID: categoriesMap["Seng"].ID,
			PricePerKg:     6500.00,
		},
		{
			ItemName:       "Seng Campur",
			ItemCategoryID: categoriesMap["Seng"].ID,
			PricePerKg:     6000.00,
		},
		// Timah
		{
			ItemName:       "Timah Batangan",
			ItemCategoryID: categoriesMap["Timah"].ID,
			PricePerKg:     70000.00,
		},
		{
			ItemName:       "Timah Gulungan",
			ItemCategoryID: categoriesMap["Timah"].ID,
			PricePerKg:     72000.00,
		},
		// Akunium
		{
			ItemName:       "Akunium Pipa",
			ItemCategoryID: categoriesMap["Akunium"].ID,
			PricePerKg:     9300.00,
		},
		{
			ItemName:       "Akunium Plat",
			ItemCategoryID: categoriesMap["Akunium"].ID,
			PricePerKg:     9500.00,
		},
		// Kaleng
		{
			ItemName:       "Kaleng Minuman",
			ItemCategoryID: categoriesMap["Kaleng"].ID,
			PricePerKg:     2500.00,
		},
		{
			ItemName:       "Kaleng Cat",
			ItemCategoryID: categoriesMap["Kaleng"].ID,
			PricePerKg:     2200.00,
		},
		// Kabel
		{
			ItemName:       "Kabel Kupas",
			ItemCategoryID: categoriesMap["Kabel"].ID,
			PricePerKg:     18000.00,
		},
		{
			ItemName:       "Timah Super",
			ItemCategoryID: categoriesMap["Timah"].ID,
			PricePerKg:     40000.00,
		},
	}

	for _, item := range items {
		if err := config.DB.Create(&item).Error; err != nil {
			log.Printf("Failed to create item %s: %v", item.ItemName, err)
		}
	}
	log.Println("✅ Items seeded successfully (19 items)")
}
