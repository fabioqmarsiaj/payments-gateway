import { Injectable } from '@nestjs/common';
import { CreateOrderDto } from './dto/create-order.dto';
import { UpdateOrderDto } from './dto/update-order.dto';
import { Order } from './entities/order.entity';
import { InjectModel } from '@nestjs/sequelize';

@Injectable()
export class OrdersService {
  constructor(@InjectModel(Order) private orderModule: typeof Order) {}

  create(createOrderDto: CreateOrderDto) {
    return this.orderModule.create({ ...createOrderDto });
  }

  findAll() {
    return this.orderModule.findAll();
  }

  findOne(id: string) {
    return this.orderModule.findByPk(id);
  }

  async update(id: string, updateOrderDto: UpdateOrderDto) {
    const order = await this.findOne(id);
    return order.update(updateOrderDto);
  }

  remove(id: string) {
    return `This action removes a #${id} order`;
  }
}
