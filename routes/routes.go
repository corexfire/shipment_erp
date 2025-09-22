package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes all routes
func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "landing.html", gin.H{
			"title": "Home",
		})
	})
	r.GET("/features", func(c *gin.Context) {
		c.HTML(http.StatusOK, "features.html", gin.H{
			"title": "Features",
		})
	})
	r.GET("/our-team", func(c *gin.Context) {
		c.HTML(http.StatusOK, "our-team.html", gin.H{
			"title": "Our Team",
		})
	})
	r.GET("/faqs", func(c *gin.Context) {
		c.HTML(http.StatusOK, "faqs.html", gin.H{
			"title": "FAQ's",
		})
	})
	r.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", gin.H{
			"title": "Contact",
		})
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Login",
		})
	})
	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "Register",
		})
	})
	r.GET("/reset-password", func(c *gin.Context) {
		c.HTML(http.StatusOK, "reset-password.html", gin.H{
			"title": "Reset Password",
		})
	})
	r.GET("/forget-password", func(c *gin.Context) {
		c.HTML(http.StatusOK, "forget-password.html", gin.H{
			"title": "Forget Password",
		})
	})
	r.GET("/confirm-mail", func(c *gin.Context) {
		c.HTML(http.StatusOK, "confirm-mail.html", gin.H{
			"title": "Confirm Mail",
		})
	})
	r.GET("/lock-screen", func(c *gin.Context) {
		c.HTML(http.StatusOK, "lock-screen.html", gin.H{
			"title": "Lock Screen",
		})
	})
	r.GET("/logout", func(c *gin.Context) {
		c.HTML(http.StatusOK, "logout.html", gin.H{
			"title": "Logout",
		})
	})

	// Dashboard Routes
	dashboard := r.Group("/dashboard")
	{
		dashboard.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "ecommarce.html", gin.H{
				"title": "eCommarce",
			})
		})
		dashboard.GET("/crm", func(c *gin.Context) {
			c.HTML(http.StatusOK, "crm.html", gin.H{
				"title": "CRM",
			})
		})
		dashboard.GET("/project-management", func(c *gin.Context) {
			c.HTML(http.StatusOK, "project-management.html", gin.H{
				"title": "Project Management",
			})
		})
		dashboard.GET("/lms", func(c *gin.Context) {
			c.HTML(http.StatusOK, "lms.html", gin.H{
				"title": "LMS",
			})
		})
		dashboard.GET("/help-desk", func(c *gin.Context) {
			c.HTML(http.StatusOK, "help-desk.html", gin.H{
				"title": "Help Desk",
			})
		})
		dashboard.GET("/analytics", func(c *gin.Context) {
			c.HTML(http.StatusOK, "analytics.html", gin.H{
				"title": "Analytics",
			})
		})
		dashboard.GET("/crypto", func(c *gin.Context) {
			c.HTML(http.StatusOK, "crypto.html", gin.H{
				"title": "Crypto",
			})
		})
		dashboard.GET("/sales", func(c *gin.Context) {
			c.HTML(http.StatusOK, "sales.html", gin.H{
				"title": "Sales",
			})
		})
		dashboard.GET("/hospital", func(c *gin.Context) {
			c.HTML(http.StatusOK, "hospital.html", gin.H{
				"title": "Hospital",
			})
		})
		dashboard.GET("/marketing", func(c *gin.Context) {
			c.HTML(http.StatusOK, "marketing.html", gin.H{
				"title": "Marketing",
			})
		})
		dashboard.GET("/nft", func(c *gin.Context) {
			c.HTML(http.StatusOK, "nft.html", gin.H{
				"title": "NFT",
			})
		})
		dashboard.GET("/saas", func(c *gin.Context) {
			c.HTML(http.StatusOK, "saas.html", gin.H{
				"title": "SaaS",
			})
		})
		dashboard.GET("/real-estate", func(c *gin.Context) {
			c.HTML(http.StatusOK, "real-estate.html", gin.H{
				"title": "Real Estate",
			})
		})
		dashboard.GET("/shipment", func(c *gin.Context) {
			c.HTML(http.StatusOK, "shipment.html", gin.H{
				"title": "Shipment",
			})
		})
		dashboard.GET("/finance", func(c *gin.Context) {
			c.HTML(http.StatusOK, "finance.html", gin.H{
				"title": "Finance",
			})
		})
		dashboard.GET("/hrm", func(c *gin.Context) {
			c.HTML(http.StatusOK, "hrm.html", gin.H{
				"title": "HRM",
			})
		})
		dashboard.GET("/school", func(c *gin.Context) {
			c.HTML(http.StatusOK, "school.html", gin.H{
				"title": "School",
			})
		})
		dashboard.GET("/call-center", func(c *gin.Context) {
			c.HTML(http.StatusOK, "call-center.html", gin.H{
				"title": "Call Center",
			})
		})
		dashboard.GET("/pos-system", func(c *gin.Context) {
			c.HTML(http.StatusOK, "pos-system.html", gin.H{
				"title": "Pos System",
			})
		})
		dashboard.GET("/podcast", func(c *gin.Context) {
			c.HTML(http.StatusOK, "podcast.html", gin.H{
				"title": "podcast",
			})
		})

		dashboard.GET("/doctor", func(c *gin.Context) {
			c.HTML(http.StatusOK, "doctor.html", gin.H{
				"title": "doctor",
			})
		})

		dashboard.GET("/social-media", func(c *gin.Context) {
			c.HTML(http.StatusOK, "social-media.html", gin.H{
				"title": "social-media",
			})
		})

		dashboard.GET("/beauty-salon", func(c *gin.Context) {
			c.HTML(http.StatusOK, "beauty-salon.html", gin.H{
				"title": "beauty-salon",
			})
		})

		dashboard.GET("/store-analytics", func(c *gin.Context) {
			c.HTML(http.StatusOK, "store-analytics.html", gin.H{
				"title": "store-analytics",
			})
		})

		dashboard.GET("/restaurant", func(c *gin.Context) {
			c.HTML(http.StatusOK, "restaurant.html", gin.H{
				"title": "restaurant",
			})
		})

		dashboard.GET("/hotel", func(c *gin.Context) {
			c.HTML(http.StatusOK, "hotel.html", gin.H{
				"title": "hotel",
			})
		})

		dashboard.GET("/real-estate-agent", func(c *gin.Context) {
			c.HTML(http.StatusOK, "real-estate-agent.html", gin.H{
				"title": "real-estate-agent",
			})
		})

		dashboard.GET("/credit-card", func(c *gin.Context) {
			c.HTML(http.StatusOK, "credit-card.html", gin.H{
				"title": "credit-card",
			})
		})

		dashboard.GET("/crypto-trader", func(c *gin.Context) {
			c.HTML(http.StatusOK, "crypto-trader.html", gin.H{
				"title": "crypto-trader",
			})
		})

		dashboard.GET("/crypto-performance", func(c *gin.Context) {
			c.HTML(http.StatusOK, "crypto-performance.html", gin.H{
				"title": "crypto-performance",
			})
		})

		dashboard.GET("/to-do-list", func(c *gin.Context) {
			c.HTML(http.StatusOK, "to-do-list.html", gin.H{
				"title": "to-do-list",
			})
		})

		dashboard.GET("/calendar", func(c *gin.Context) {
			c.HTML(http.StatusOK, "calendar.html", gin.H{
				"title": "calendar",
			})
		})

		dashboard.GET("/contacts", func(c *gin.Context) {
			c.HTML(http.StatusOK, "contacts.html", gin.H{
				"title": "contacts",
			})
		})

		dashboard.GET("/chat", func(c *gin.Context) {
			c.HTML(http.StatusOK, "chat.html", gin.H{
				"title": "chat",
			})
		})

		dashboard.GET("/inbox", func(c *gin.Context) {
			c.HTML(http.StatusOK, "inbox.html", gin.H{
				"title": "inbox",
			})
		})

		dashboard.GET("/compose", func(c *gin.Context) {
			c.HTML(http.StatusOK, "compose.html", gin.H{
				"title": "compose",
			})
		})

		dashboard.GET("/read-email", func(c *gin.Context) {
			c.HTML(http.StatusOK, "read-email.html", gin.H{
				"title": "read-email",
			})
		})
		dashboard.GET("/snoozed", func(c *gin.Context) {
			c.HTML(http.StatusOK, "snoozed.html", gin.H{
				"title": "snoozed",
			})
		})
		dashboard.GET("/draft", func(c *gin.Context) {
			c.HTML(http.StatusOK, "draft.html", gin.H{
				"title": "draft",
			})
		})
		dashboard.GET("/sent-mail", func(c *gin.Context) {
			c.HTML(http.StatusOK, "sent-mail.html", gin.H{
				"title": "sent-mail",
			})
		})
		dashboard.GET("/trash-email", func(c *gin.Context) {
			c.HTML(http.StatusOK, "trash-email.html", gin.H{
				"title": "trash-email",
			})
		})
		dashboard.GET("/spam", func(c *gin.Context) {
			c.HTML(http.StatusOK, "spam.html", gin.H{
				"title": "spam",
			})
		})
		dashboard.GET("/starred", func(c *gin.Context) {
			c.HTML(http.StatusOK, "starred.html", gin.H{
				"title": "starred",
			})
		})
		dashboard.GET("/important", func(c *gin.Context) {
			c.HTML(http.StatusOK, "important.html", gin.H{
				"title": "important",
			})
		})
		dashboard.GET("/orders", func(c *gin.Context) {
			c.HTML(http.StatusOK, "shipment-order.html", gin.H{
				"title": "Shipment Orders List",
			})
		})
		dashboard.GET("/orders/create", func(c *gin.Context) {
			c.HTML(http.StatusOK, "shipment-order-create.html", gin.H{
				"title": "New Orders",
			})
		})
		dashboard.GET("/orders/pending", func(c *gin.Context) {
			c.HTML(http.StatusOK, "shipment-order-pending.html", gin.H{
				"title": "Pending Orders",
			})
		})
		dashboard.GET("/orders/completed", func(c *gin.Context) {
			c.HTML(http.StatusOK, "shipment-order-completed.html", gin.H{
				"title": "completed Orders",
			})
		})
		//documents
		dashboard.GET("/documents", func(c *gin.Context) {
			c.HTML(http.StatusOK, "shipment-documents.html", gin.H{
				"title": "Document List",
			})
		})
		dashboard.GET("/documents/upload", func(c *gin.Context) {
			c.HTML(http.StatusOK, "shipment-documents-upload.html", gin.H{
				"title": "Document Upload",
			})
		})
		dashboard.GET("/documents/ci", func(c *gin.Context) {
			c.HTML(http.StatusOK, "shipment-documents-ci.html", gin.H{
				"title": "Commercial Invoice",
			})
		})
		dashboard.GET("/documents/pl", func(c *gin.Context) {
			c.HTML(http.StatusOK, "shipment-documents-pl.html", gin.H{
				"title": "Packing List",
			})
		})
		dashboard.GET("/documents/bl", func(c *gin.Context) {
			c.HTML(http.StatusOK, "shipment-documents-bl.html", gin.H{
				"title": "Bill of Lading",
			})
		})
		dashboard.GET("/documents/ls", func(c *gin.Context) {
			c.HTML(http.StatusOK, "shipment-documents-ls.html", gin.H{
				"title": "Laporan Surveyor",
			})
		})
		dashboard.GET("/documents/fe", func(c *gin.Context) {
			c.HTML(http.StatusOK, "shipment-documents-fe.html", gin.H{
				"title": "Form E",
			})
		})
	}

}
