# %%
import json
import os

from creatorsapi_python_sdk.api_client import ApiClient
from creatorsapi_python_sdk.api.default_api import DefaultApi
from creatorsapi_python_sdk.models.search_items_request_content import SearchItemsRequestContent
from creatorsapi_python_sdk.exceptions import ApiException

from rec_book_client import Book

import dotenv
dotenv.load_dotenv()  # Carrega as variáveis de ambiente do arquivo .env

AMAZON_ID_CREDENCIAL = os.getenv("AMAZON_ID_CREDENCIAL")
AMAZON_SECRET = os.getenv("AMAZON_SECRET")
AMAZON_VERSION = os.getenv("AMAZON_VERSION")
AMAZON_TAG = os.getenv("AMAZON_TAG")
AMAZON_COUNTRY = os.getenv("AMAZON_COUNTRY")


class AmazonClient:
    def __init__(self):

        self.api_client = ApiClient(
            credential_id=AMAZON_ID_CREDENCIAL,
            credential_secret=AMAZON_SECRET,
            version=AMAZON_VERSION
        )

        self.api = DefaultApi(self.api_client)

    def search_items(self, book: Book, search_index: str = "Books", sort_by: str = None):
        
        """
        Choose resources you want from SearchItemsResource enum
        For more details, refer: https://affiliate-program.amazon.com/creatorsapi/docs/en-us/api-reference/operations/search-items#resources-parameter
        """
        
        resources = [
            'images.primary.medium',
            'itemInfo.title',
            'itemInfo.productInfo',
            'offersV2.listings.availability',
            'offersV2.listings.condition',
            'offersV2.listings.dealDetails',
            'offersV2.listings.isBuyBoxWinner',
            'offersV2.listings.loyaltyPoints',
            'offersV2.listings.merchantInfo',
            'offersV2.listings.price',
            'offersV2.listings.type'
        ]
        
        # Create SearchItems request
        search_items_request = SearchItemsRequestContent(
            partner_tag=AMAZON_TAG,
            author=book.author,
            title=book.title,
            search_index=search_index,
            resources=resources,
            sortBy=sort_by,
        )
        
        try:
            response = self.api.search_items(x_marketplace="www.amazon.com.br",
                                             search_items_request_content=search_items_request)
            
            data = response.to_dict() if hasattr(response, 'to_dict') else str(response)
            return data
            
        except ApiException as exception:
            print('Error calling Creators API!')
            print(exception)
            return {"error": str(exception)}
        
        except Exception as exception:
            print('Unexpected error:', exception)
            return {"error": str(exception)}


# %%

# AvgCustomerReviews
# Featured
# NewestArrivals
# Price:HighToLow
# Price:LowToHigh
# Relevance

