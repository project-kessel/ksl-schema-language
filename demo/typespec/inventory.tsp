import "@typespec/rest";
import "@typespec/openapi";
import "./decorators.tsp";

using TypeSpec.Http;

@service({
  title: "Inventory Service",
})
@doc("This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters.")
namespace Inventory;

// Model types
model Host {
  name: string;
  tag?: string;

  @minValue(0)
  @maxValue(20)
  age: int32;
}

model Toy {
  id: int64;
  petId: int64;
  name: string;
}

@doc("Error")
@error
model Error {
  code: int32;
  message: string;
}

@doc("Not modified")
model NotModified<Body> {
  @statusCode _: 304;
  @body body: Body;
}

@friendlyName("{name}ListResults", Item)
model ResponsePage<Item> {
  items: Item[];
  nextLink?: string;
}

model PetId {
  @path petId: int32;
}

@doc("Manage your hosts.")
@route("/hosts")
namespace Hosts {
  @doc("Returns a host")
  @KesselLib.prem_read("read Host")
  op read(...PetId): Host | (NotModifiedResponse & Host) | Error;
}